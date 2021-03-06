package kmeans

import (
	"sync"
)

func kmeansWorker1(data []ClusteredObservation, mean []Observation, mLen []int, meanLockers []sync.Mutex, done chan<- int) {
	for _, v := range data {
		num := v.ClusterNumber
		meanLockers[num].Lock()
		mean[num].Add(v.Observation)
		mLen[num]++
		meanLockers[num].Unlock()
	}
	done <- 0
}

func kmeansWorker2(data []ClusteredObservation, mean []Observation, done chan<- int) {
	changes := 0
	for i, v := range data {
		if closestCluster, _ := Near(v, mean, EuclideanDistance); closestCluster != v.ClusterNumber {
			data[i].ClusterNumber = closestCluster
			changes++
		}
	}
	done <- changes
}
