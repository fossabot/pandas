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

import "encoding/json"

const (
	JOB_STATUS_CREATED = "created"
	JOB_STATUS_KILLED  = "killed"
)

type Job struct {
	ID      string `json:"id"`
	Domain  string `json:"domain"`  // job's domain
	Payload []byte `json:"payload"` // job's payload
	Status  string `json:"status"`  // job's  status
}

func NewJob() *Job {
	return &Job{Payload: []byte{}}
}

func (job *Job) MarshalBinary() ([]byte, error) {
	return json.Marshal(job)
}

func (job *Job) UnmarshalBinary(buf []byte) error {
	return json.Unmarshal(buf, job)
}
