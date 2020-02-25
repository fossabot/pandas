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

import "github.com/sirupsen/logrus"

const (
	HEADMAST_SCHED_POLICY_ROUNDBIN = "roundbin"
	HEADMAST_SCHED_POLICY_BLANCE   = "balance"
)

// SchedulePolicy determain scheduler policy that affect how the new job and
// worker should be scheduled
// SchedulePolicy doesn't update etcd directly, it just caculated the result
type SchedulePolicy interface {

	// DetermainWithJobChange caculate the last result for scheduler when jobs
	// is added or deleted
	DetermainWithJobChanged(ctx *ScheduleContext, affectedJob *Job, allWorkers []*Worker, added bool) []*Worker

	// DetermainWithWorkerChange caculate the last result for scheduler when
	// worker is added or deleted
	DetermainWithWorkerChanged(ctx *ScheduleContext, affectedWorker *Worker, allWorkers []*Worker, added bool) []*Worker
}

// NewSchedulePolicy return scheduler policy based on serving options's
// schedule-policy
func NewSchedulePolicy(servingOptions *ServingOptions) SchedulePolicy {
	switch servingOptions.SchedulePolicy {
	default:
		return &roundbinPolicy{}
	}
	return nil
}

// roundbinPolicy is default scheduler policy
type roundbinPolicy struct{}

func (r *roundbinPolicy) DetermainWithWorkerChanged(ctx *ScheduleContext, affectedWorker *Worker, allWorkers []*Worker, added bool) []*Worker {
	workers := []*Worker{}

	if added {
		// when worker changed notification is received, the worker had register
		// itself in '/headmast/workers', as for roudbin policy, there are
		// nothting to do with the notification.
		logrus.Infof("new worker '%s' had beed added", affectedWorker.ID)
	} else {
		// when worker is removed from '/headmast/workers', all jobs of the
		// worker should be assigned to other workers.  all workers contain the
		// affected worker
		leftWorkers := []*Worker{}
		for _, worker := range allWorkers {
			if worker.ID == affectedWorker.ID {
				leftWorkers = append(leftWorkers, worker)
				break
			}
		}
		if len(leftWorkers) == 0 {
			if len(affectedWorker.WorkingJobs) > 0 {
				logrus.Info("no available workers, but working jobs still exist")
				return workers
			}
			logrus.Infof("no availabel workers and jobs existed now")
			return workers
		}
		// Assign left jobs to other workers
		index := 0
		affectedWorkers := make(map[string]*Worker)
		for _, jobID := range affectedWorker.WorkingJobs {
			if index >= len(leftWorkers) {
				index = 0
			}
			worker := leftWorkers[index]
			worker.WorkingJobs = append(worker.WorkingJobs, jobID)
			affectedWorkers[worker.ID] = worker
			index++
		}
		for _, worker := range affectedWorkers {
			workers = append(workers, worker)
		}
	}
	return workers
}

func (r *roundbinPolicy) DetermainWithJobChanged(ctx *ScheduleContext, affectedJob *Job, allWorkers []*Worker, added bool) []*Worker {
	workers := []*Worker{}

	if added {
		// The new job should be assigned to last scheduled worker based on context
		if len(allWorkers) == 0 {
			logrus.Warningf("no workers to schedule the job '%s'", affectedJob.ID)
			return workers
		}
		if ctx.LastScheduledWorkerIndex > len(allWorkers) {
			ctx.LastScheduledWorkerIndex = 0
		}
		worker := allWorkers[ctx.LastScheduledWorkerIndex]
		ctx.LastScheduledWorkerIndex++
		worker.WorkingJobs = append(worker.WorkingJobs, affectedJob.ID)
		workers = append(workers, worker)
	} else {
		// The deleted job should be remove from worker's working path and
		// moved to killing path
		for _, worker := range allWorkers {
			for index, jobID := range worker.WorkingJobs {
				// If the deleted job is found in working path, remove it
				if jobID == affectedJob.ID {
					worker.WorkingJobs = append(worker.WorkingJobs[:index], worker.WorkingJobs[index:]...)
					workers = append(workers, worker)
					return workers
				}
			}
		}
	}
	return workers
}

// balancePolicy will reassign worker's jobs based on worker's load
type balancePolicy struct{}
