package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Adapter contains information about the adapter.
type Adapter struct {
	Name    string `json:"name" bson:"name"`
	Type    string `json:"type" bson:"type"`
	Version string `json:"version" bson:"version"`
}

// Logs contains all properties of a log.
type Logs struct {
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
	Type      string `json:"type" bson:"type"`
	Log       string `json:"log" bson:"log"`
}

// Metrics contains information about the system
type Metrics struct {
	Platform string `json:"platform,omitempty" bson:"platform,omitempty"`
	Browser  string `json:"browser,omitempty" bson:"browser,omitempty"`
	IsMobile string `json:"isMobile,omitempty" bson:"isMobile,omitempty"`
}

// UserInteraction contains information about an element that was clicked by the user.
type UserInteraction struct {
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
	Element   string `json:"element" bson:"element"`
	InnerText string `json:"innerText" bson:"innerText"`
	ElementID string `json:"elementId" bson:"elementId"`
	Location  string `json:"location" bson:"location"`
}

// Error contains all properties of an error event.
type Error struct {
	ID               *primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Message          string               `json:"message" bson:"message"`
	Stacktrace       string               `json:"stacktrace" bson:"stacktrace"`
	Evolution        map[string]int       `json:"evolution" bson:"evolution"`
	Path             string               `json:"path" bson:"path"`
	Line             string               `json:"line" bson:"line"`
	Type             string               `json:"type" bson:"type"`
	Adapter          Adapter              `json:"adapter" bson:"adapter"`
	Fingerprint      string               `json:"fingerprint" bson:"fingerprint"`
	Badges           map[string]string    `json:"badges,omitempty" bson:"badges,omitempty"`
	Snippet          map[string]string    `json:"snippet,omitempty" bson:"snippet,omitempty"`
	Logs             []Logs               `json:"logs,omitempty" bson:"logs,omitempty"`
	Ticket           string               `json:"ticket" bson:"ticket"`
	Host             string               `json:"host,omitempty" bson:"host,omitempty"`
	UserAgent        string               `json:"userAgent" bson:"userAgent"`
	Metrics          Metrics              `json:"metrics" bson:"metrics"`
	UserInteractions []UserInteraction    `json:"userInteractions,omitempty" bson:"userInteractions,omitempty"`
	AnonymizeData    bool                 `json:"anonymizeData" bson:"-"`
	ClientIP         string               `json:"clientIp" bson:"clientIp"`
	Count            int                  `json:"count,omitempty" bson:"count,omitempty"`
	Timestamp        int64                `json:"timestamp" bson:"timestamp"`
	Resolved         bool                 `json:"resolved" bson:"resolved"`
	SeenBy           []primitive.ObjectID `json:"seenBy,omitempty" bson:"seenBy,omitempty"`
	LastSeen         int64                `json:"lastSeen" bson:"lastSeen"`
	CreatedAt        time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time            `json:"updatedAt" bson:"updatedAt"`
}

// IsValid validates an error event to make sure the data is not too large.
func (e *Error) IsValid() bool {
	if len(e.Logs) > 50 {
		return false
	}

	if len(e.UserInteractions) > 50 {
		return false
	}

	if len(e.SeenBy) > 0 {
		return false
	}

	if len(e.Badges) > 200 {
		return false
	}

	if len(e.Snippet) > 50 {
		return false
	}

	if len(e.Evolution) > 0 {
		return false
	}

	return true
}

// AnalyticData contains aggregated analytic data.
type AnalyticData struct {
	Day             int64          `json:"day" bson:"day"`
	Hour            int64          `json:"hour" bson:"hour"`
	Windows         int            `json:"windows,omitempty" bson:"w,omitempty"`
	Mac             int            `json:"mc,omitempty" bson:"m,omitempty"`
	Linux           int            `json:"linux,omitempty" bson:"l,omitempty"`
	OtherPlatforms  int            `json:"otherPlatforms,omitempty" bson:"oP,omitempty"`
	Chrome          int            `json:"chrome,omitempty" bson:"c,omitempty"`
	Firefox         int            `json:"firefox,omitempty" bson:"f,omitempty"`
	Safari          int            `json:"safari,omitempty" bson:"s,omitempty"`
	Edge            int            `json:"edge,omitempty" bson:"e,omitempty"`
	IE              int            `json:"ie,omitempty" bson:"i,omitempty"`
	Opera           int            `json:"opera,omitempty" bson:"o,omitempty"`
	OtherBrowsers   int            `json:"otherBrowsers,omitempty" bson:"oB,omitempty"`
	Mobile          int            `json:"mobile,omitempty" bson:"mbl,omitempty"`
	Tablet          int            `json:"tablet,omitempty" bson:"t,omitempty"`
	Desktop         int            `json:"desktop,omitempty" bson:"d,omitempty"`
	Visits          int            `json:"visits,omitempty" bson:"v,omitempty"`
	NewVisitors     int            `json:"newVisitors,omitempty" bson:"n,omitempty"`
	TotalSessions   int            `json:"sessions,omitempty" bson:"tS,omitempty"`
	TotalTimeOnPage int            `json:"totalTimeOnPage,omitempty" bson:"tT,omitempty"`
	Pages           map[string]int `json:"pages,omitempty" bson:"p,omitempty"`
	Referrer        map[string]int `json:"referrer,omitempty" bson:"r,omitempty"`
}

// Analytics represents a ressource in the database.
type Analytics struct {
	Ticket                string                  `json:"ticket" bson:"ticket"`
	Month                 int64                   `json:"month" bson:"month"`
	HumanReadableMonth    string                  `json:"humanReadableMonth" bson:"humanReadableMonth"`
	AggregatedMonthlyData AnalyticData            `json:"aggregatedMonthlyData" bson:"aggregatedMonthlyData"`
	Data                  map[string]AnalyticData `json:"data" bson:"data"`
	CreatedAt             time.Time               `json:"createdAt" bson:"createdAt"`
	UpdatedAt             time.Time               `json:"updatedAt" bson:"updatedAt"`
}

// AnalyticEvent contains information about a page visitor sent by the adapter.
type AnalyticEvent struct {
	Ticket       string `json:"ticket" bson:"ticket"`
	IsNewVisitor bool   `json:"isNewVisitor" bson:"isNewVisitor"`
	IsNewSession bool   `json:"isNewSession" bson:"isNewSession"`
	TimeOnPage   int    `json:"timeOnPage" bson:"timeOnPage"`
	Referrer     string `json:"referrer" bson:"referrer"`
	Page         string `json:"page" bson:"page"`
	UserAgent    string `json:"userAgent" bson:"userAgent"`
}

// AnalyticInsights containes aggeregated analytic data for a specific time frame.
type AnalyticInsights struct {
	TimeframeStart   int64          `json:"timeframeStart"`
	TimeframeEnd     int64          `json:"timeframeEnd"`
	TotalVisits      int            `json:"totalVisits"`
	TotalNewVisitors int            `json:"totalNewVisitors"`
	TotalSessions    int            `json:"totalSessions"`
	Data             []AnalyticData `json:"pageViews"`
}
