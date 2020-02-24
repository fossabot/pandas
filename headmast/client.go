//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.

package headmast

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type ClientOptions struct {
	ServerAddr string // header master server address
}

// Client is endpoint for headmast
type Client struct {
	options *ClientOptions
}

// NewClient create a client endpoint
func NewClient(opts *ClientOptions) *Client {
	return &Client{options: opts}
}

// CreateNewJob create a new job in headmast
func (c *Client) CreateNewJob(job *Job) error {
	body, err := json.Marshal(job)
	if err != nil {
		logrus.WithError(err)
		return err
	}

	url := fmt.Sprintf("http://%s/api/v1/jobs", c.options.ServerAddr)
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		logrus.WithError(err)
		return err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err)
	}
	return err
}

// DeleteJob remove a job from headmast
func (c *Client) DeleteJob(jobID string) error {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s/api/v1/jobs/%s", c.options.ServerAddr, jobID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		logrus.WithError(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err)
	}
	return err
}

// GetJob return job detail from headmast
func (c *Client) GetJob(jobID string) (*Job, error) {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s/api/v1/jobs/%s", c.options.ServerAddr, jobID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.WithError(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err)
		return nil, err
	}
	job := &Job{}
	if err := json.Unmarshal(body, job); err != nil {
		logrus.WithError(err)
		return nil, err
	}
	return job, nil
}

// GetJobsWithDomain return jobs within specific domain
func (c *Client) GetJobsWithDomain(domain string) ([]*Job, error) {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s/api/v1/jobs", c.options.ServerAddr)
	req, err := http.NewRequest("GET", url, strings.NewReader(fmt.Sprintf("domain=%s", domain)))
	if err != nil {
		logrus.WithError(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err)
		return nil, err
	}
	jobs := []*Job{}
	if err := json.Unmarshal(body, jobs); err != nil {
		logrus.WithError(err)
		return nil, err
	}
	return jobs, nil
}

// GetJobs return current all jobs
func (c *Client) GetJobs() ([]*Job, error) {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s/api/v1/jobs", c.options.ServerAddr)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.WithError(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err)
		return nil, err
	}
	jobs := []*Job{}
	if err := json.Unmarshal(body, jobs); err != nil {
		logrus.WithError(err)
		return nil, err
	}
	return jobs, nil
}

type WatchPathHandler func(jobPath string, payload []byte)

// WatchJobPath watch a specific job path and call handler when envent occured
func (c *Client) WatchJobPath(jobPath string, handler WatchPathHandler) error {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s/api/v1/watch", c.options.ServerAddr)
	req, err := http.NewRequest("GET", url, strings.NewReader(fmt.Sprintf("path=%s", jobPath)))
	if err != nil {
		logrus.WithError(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	errCh := make(chan error)
	go func(path string, handler WatchPathHandler, reader io.Reader, errCh chan error) {
		buf := []byte{}
		for {
			n, err := reader.Read(buf)
			if err != nil {
				errCh <- err
				return
			}
			handler(path, buf[:n])
		}
	}(jobPath, handler, resp.Body, errCh)
	<-errCh
	return nil
}

// ControlJob control a job to be killed or alived
func (c *Client) ControlJob(jobID string, action string) {
	url := fmt.Sprintf("http://%s/api/v1/jobs/%s", c.options.ServerAddr, jobID)
	resp, err := http.Post(url,
		"application/json",
		strings.NewReader(fmt.Sprintf("action=%s", action)),
	)
	if err != nil {
		logrus.WithError(err)
		return
	}
	defer resp.Body.Close()
}
