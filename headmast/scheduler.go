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

// ScheduleContext encapsulate informations for job sechduling
type ScheduleContext struct {
	LastScheduledWorkerIndex int
}

// NewScheduleContex create default schedule context
func NewScheduleContext() *ScheduleContext {
	return &ScheduleContext{
		LastScheduledWorkerIndex: 0,
	}
}

// JobScheduler schedule job based on available workers
type JobScheduler interface {
}

// NewJobSchduler return instance of job scheduler that based on schedule
// policy created from serving options
func NewJobScheduler(servingOptions *ServingOptions, jobManager JobManager, workerManager WorkerManager) JobScheduler {
	return newJobScheduler(servingOptions, jobManager, workerManager)
}

// jobScheduler is default implementations of job scheduler
type jobScheduler struct {
	servingOptions *ServingOptions
	context        *ScheduleContext
	jobManager     JobManager
	workerManager  WorkerManager
	schedulePolicy SchedulePolicy
}

// newJobScheduler return default job scheduler instance
func newJobScheduler(servingOptions *ServingOptions, jobManager JobManager, workerManager WorkerManager) JobScheduler {
	s := &jobScheduler{
		servingOptions: servingOptions,
		context:        NewScheduleContext(),
		jobManager:     jobManager,
		workerManager:  workerManager,
		schedulePolicy: NewSchedulePolicy(servingOptions),
	}
	workerManager.RegisterObserver(s.onWorkerChanges)
	jobManager.RegisterObserver(s.onJobChanges)
	return s
}

// onWorkerChanges is called when worker nodes are added or removed
func (s *jobScheduler) onWorkerChanges(w *Worker, reason string) {
	switch reason {
	case HEADMAST_CHANGES_ADDED:
		// when worker is added, we should reblance all jobs based on all
		// workers and scheduler policy
		workers := s.workerManager.GetWorkers()
		affectedWorkers := s.schedulePolicy.DeterminWithWorkerChanged(s.context, w, workers, true)
		s.workerManager.UpdateWorkers(affectedWorkers)

	case HEADMAST_CHANGES_DELETED:
		// when worker is deleted, get jobs of the deleted workers and reassign
		// them to other workers
		workers := s.workerManager.GetWorkers()
		affectedWorkers := s.schedulePolicy.DeterminWithWorkerChanged(s.context, w, workers, false)
		s.workerManager.UpdateWorkers(affectedWorkers)
	default:
	}
}

// onJobChanges is called when job is added or removed
func (s *jobScheduler) onJobChanges(job *Job, reason string) {
	switch reason {
	case HEADMAST_CHANGES_ADDED:
		// The job is already placed on '/headmast/jobs' by client request,
		// the scheduler should place the job on specific worker path based on
		// schedule policy
		workers := s.workerManager.GetWorkers()
		affectedWorkers := s.schedulePolicy.DeterminWithJobChanged(s.context, job, workers, true)
		s.workerManager.UpdateWorkers(affectedWorkers)

	case HEADMAST_CHANGES_DELETED:
		workers := s.workerManager.GetWorkers()
		affectedWorkers := s.schedulePolicy.DeterminWithJobChanged(s.context, job, workers, false)
		s.workerManager.UpdateWorkers(affectedWorkers)
	default:
	}

}
