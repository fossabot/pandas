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
package raft

import (
	"bytes"
	"context"
	"log"
	"strings"
	"sync"

	"github.com/cloustone/pandas/pkg/service"
	logr "github.com/sirupsen/logrus"
	"go.etcd.io/etcd/etcdserver/api/snap"
	"go.etcd.io/etcd/raft/raftpb"
	"golang.org/x/sync/errgroup"
)

const (
	SERVICE_NAME = "synchronizer"
)

type syncService struct {
	context          context.Context
	shutdownFn       context.CancelFunc
	childRoutines    *errgroup.Group
	dataStore        Repository
	mutex            sync.RWMutex
	proposeC         chan string // channel for proposing updates
	snapshotter      *snap.Snapshotter
	commitC          <-chan *string
	errorC           <-chan error
	snapshotterReady <-chan *snap.Snapshotter
	confChangeC      <-chan raftpb.ConfChange
}

func init() {
	service.Add(1, createDataService)
}

func createDataService() (service.Service, error) {
	rootCtx, shutdownFn := context.WithCancel(context.Background())
	childRoutines, childCtx := errgroup.WithContext(rootCtx)

	s := &syncService{
		context:       childCtx,
		shutdownFn:    shutdownFn,
		childRoutines: childRoutines,
		mutex:         sync.RWMutex{},
		proposeC:      make(chan string),
		confChangeC:   make(chan raftpb.ConfChange),
	}
	// TODO: how to get the following parameter
	var id int
	var cluster string
	var join bool

	// raft provides a commit stream for the proposals from the http api
	s.commitC, s.errorC, s.snapshotterReady = newRaftNode(id, strings.Split(cluster, ","), join, s.getSnapshot, s.proposeC, s.confChangeC)

	// replay log into key-value map
	s.readCommits(s.commitC, s.errorC)
	// read commits from raft into kvStore map until error
	go s.readCommits(s.commitC, s.errorC)
	return s, nil
}

func (s *syncService) Start() error {
	return nil
}

func (s *syncService) Name() string { return SERVICE_NAME }

func (s *syncService) Shutdown() {
	s.shutdownFn()

	err := s.childRoutines.Wait()
	if err != nil && err != context.Canceled {
		logr.WithError(err).Errorf("data service shutdown failed")
	}
}

func (s *syncService) Propose(k string, v string) {
	var buf bytes.Buffer
	/*
		if err := gob.NewEncoder(&buf).Encode(kv{k, v}); err != nil {
			log.Fatal(err)
		}
	*/
	s.proposeC <- buf.String()
}

func (s *syncService) Lookup(key string) (string, bool) {
	/*
		s.mutex.RLock()
		defer s.mutex.RUnlock()
		return v, ok
	*/
	return "", false
}

func (s *syncService) readCommits(commitC <-chan *string, errorC <-chan error) {
	for data := range commitC {
		if data == nil {
			// done replaying log; new data incoming
			// OR signaled to load snapshot
			snapshot, err := s.snapshotter.Load()
			if err == snap.ErrNoSnapshot {
				return
			}
			if err != nil {
				log.Panic(err)
			}
			log.Printf("loading snapshot at term %d and index %d", snapshot.Metadata.Term, snapshot.Metadata.Index)
			if err := s.recoverFromSnapshot(snapshot.Data); err != nil {
				log.Panic(err)
			}
			continue
		}
		/*

			var dataKv kv
			dec := gob.NewDecoder(bytes.NewBufferString(*data))
			if err := dec.Decode(&dataKv); err != nil {
				log.Fatalf("raftexample: could not decode message (%v)", err)
			}
			s.mu.Lock()
			s.kvStore[dataKv.Key] = dataKv.Val
			s.mu.Unlock()
		*/
	}
	if err, ok := <-errorC; ok {
		log.Fatal(err)
	}
}

func (s *syncService) getSnapshot() ([]byte, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return []byte{}, nil
}

func (s *syncService) recoverFromSnapshot(snapshot []byte) error {
	/*
		var store map[string]string
		if err := json.Unmarshal(snapshot, &store); err != nil {
			return err
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
	*/
	return nil
}
