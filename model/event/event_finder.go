package event

import (
	"github.com/evergreen-ci/evergreen/db"
	"gopkg.in/mgo.v2/bson"
)

// === DB Logic ===

func Find(query db.Q) ([]Event, error) {
	events := []Event{}
	err := db.FindAllQ(Collection, query, &events)
	return events, err
}

// === Queries ===

// Host Events
func HostEventsForId(id string) db.Q {
	return db.Query(bson.D{
		{DataKey + "." + ResourceTypeKey, ResourceTypeHost},
		{ResourceIdKey, id},
	})
}

func MostRecentHostEvents(id string, n int) db.Q {
	return HostEventsForId(id).Sort([]string{"-" + TimestampKey}).Limit(n)
}

func HostEventsInOrder(id string) db.Q {
	return HostEventsForId(id).Sort([]string{TimestampKey})
}

// Task Events
func TaskEventsForId(id string) db.Q {
	return db.Query(bson.D{
		{DataKey + "." + ResourceTypeKey, ResourceTypeTask},
		{ResourceIdKey, id},
	})
}

func MostRecentTaskEvents(id string, n int) db.Q {
	return TaskEventsForId(id).Sort([]string{"-" + TimestampKey}).Limit(n)
}

func TaskEventsInOrder(id string) db.Q {
	return TaskEventsForId(id).Sort([]string{TimestampKey})
}

// Distro Events
func DistroEventsForId(id string) db.Q {
	return db.Query(bson.D{
		{DataKey + "." + ResourceTypeKey, ResourceTypeDistro},
		{ResourceIdKey, id},
	})
}

func MostRecentDistroEvents(id string, n int) db.Q {
	return DistroEventsForId(id).Sort([]string{"-" + TimestampKey}).Limit(n)
}

func DistroEventsInOrder(id string) db.Q {
	return DistroEventsForId(id).Sort([]string{TimestampKey})
}

// Scheduler Events
func SchedulerEventsForId(distroId string) db.Q {
	return db.Query(bson.D{
		{DataKey + "." + ResourceTypeKey, ResourceTypeScheduler},
		{ResourceIdKey, distroId},
	})
}

func RecentSchedulerEvents(distroId string, n int) db.Q {
	return SchedulerEventsForId(distroId).Sort([]string{"-" + TimestampKey}).Limit(n)
}
