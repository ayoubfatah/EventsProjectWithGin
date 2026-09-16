package models

import (
	"fmt"
	"gin-quickstart/db"
	"time"
)

type Event struct {
    Id            int64     `json:"id"`
    Name          string    `json:"name" binding:"required"`
    Slug          string    `json:"slug" binding:"required"`
    City          string    `json:"city" binding:"required"`
    Location      string    `json:"location" binding:"required"`
    Date          time.Time `json:"date" binding:"required"`
    OrganizerName string    `json:"organizerName" binding:"required"`
    ImageUrl      string    `json:"imageUrl" binding:"required"`
    Description   string    `json:"description" binding:"required"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
    UserId        int64     `json:"userId"`
}

type PaginatedEvents struct {
	Events  []Event `json:"events"`
	Page    int     `json:"page"`
	Limit   int     `json:"limit"`
	HasMore bool    `json:"hasMore"`
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
		INSERT INTO events(
			name,
			slug,
			city,
			location,
			date,
			organizerName,
			imageUrl,
			description,
			userId
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		e.Name,
		e.Slug,
		e.City,
		e.Location,
		e.Date,
		e.OrganizerName,
		e.ImageUrl,
		e.Description,
		e.UserId,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.Id = id

	return nil
}

func GetAllEvents(page int, limit int) (PaginatedEvents, error) {
	offset := (page - 1) * limit

	query := `
		SELECT
			id,
			name,
			slug,
			city,
			location,
			date,
			organizerName,
			imageUrl,
			description,
			userId
		FROM events
		ORDER BY id DESC
		LIMIT ? OFFSET ?
	`

	rows, err := db.DB.Query(query, limit+1, offset)
	if err != nil {
		return PaginatedEvents{}, err
	}
	defer rows.Close()

	events := make([]Event, 0, limit+1)

	for rows.Next() {

		var event Event

		err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.Slug,
			&event.City,
			&event.Location,
			&event.Date,
			&event.OrganizerName,
			&event.ImageUrl,
			&event.Description,
			&event.UserId,
		)

		if err != nil {
			return PaginatedEvents{}, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return PaginatedEvents{}, err
	}

	hasMore := len(events) > limit

	// We don't send the extra event to the frontend.
	if hasMore {
		events = events[:limit]
	}

	return PaginatedEvents{
		Events:  events,
		Page:    page,
		Limit:   limit,
		HasMore: hasMore,
	}, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var event Event

	err = row.Scan(
		&event.Id,
		&event.Name,
		&event.Slug,
		&event.City,
		&event.Location,
		&event.Date,
		&event.OrganizerName,
		&event.ImageUrl,
		&event.Description,
		&event.UserId,
	)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
func GetEventsByCity(city string) ([]Event, error) {
	query := `SELECT * FROM events WHERE city = ? 
			  ORDER BY id DESC
	`

	rows, err := db.DB.Query(query, city)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event

		err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.Slug,
			&event.City,
			&event.Location,
			&event.Date,
			&event.OrganizerName,
			&event.ImageUrl,
			&event.Description,
			&event.UserId,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
















func GetEventBySlug(slug string) (*Event, error) {
	query := `SELECT * FROM events WHERE slug = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(slug)

	var event Event

	err = row.Scan(
		&event.Id,
		&event.Name,
		&event.Slug,
		&event.City,
		&event.Location,
		&event.Date,
		&event.OrganizerName,
		&event.ImageUrl,
		&event.Description,
		&event.UserId,
	)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
// Update

func (e *Event) Update() error {
	query := `
		UPDATE events
		SET
			name = ?,
			slug = ?,
			city = ?,
			location = ?,
			date = ?,
			organizerName = ?,
			imageUrl = ?,
			description = ?
		WHERE id = ?
	`





	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		e.Name,
		e.Slug,
		e.City,
		e.Location,
		e.Date,
		e.OrganizerName,
		e.ImageUrl,
		e.Description,
		e.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Delete() error {
	query := `
		DELETE FROM events
		WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id)

	if err != nil {
		return err
	}

	return nil
}

// Register

func (e *Event) Register(userId int64) error {
	query := `
		INSERT INTO registrations(eventId, userId)
		VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id, userId)

	if err != nil {
		return err
	}

	return nil
}

// Cancel Registration

func (e *Event) CancelRegistration(userId int64) error {
	query := `
		DELETE FROM registrations
		WHERE eventId = ? AND userId = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id, userId)

	if err != nil {
		return err
	}

	return nil
}



func (e *Event) IsRegistered(userId int64) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM registrations
			WHERE userId = ? AND eventId = ?
		)
	`

	var registered bool

	err := db.DB.QueryRow(query, userId, e.Id).Scan(&registered)
	if err != nil {
		return false, err
	}
	fmt.Println("waaaa")
	fmt.Println(registered)
	return registered, nil
}



func GetEventsByUserID(userID int64) ([]Event , error){


query := `
    SELECT
        id,
        name,
        slug,
        city,
        location,
        date,
        organizerName,
        imageUrl,
        description,
        userId
    FROM events
    WHERE userId = ?
    ORDER BY id DESC

`

	rows , err := db.DB.Query(query, userID)
	if err != nil {
		return nil , err
	}
	defer rows.Close()

	events := make([]Event, 0)

	for rows.Next(){
		var event Event
		err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.Slug,
			&event.City,
			&event.Location,
			&event.Date,
			&event.OrganizerName,
			&event.ImageUrl,
			&event.Description,
			&event.UserId,
		)
		
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
    return nil, err
}
	return events , nil
}
 






func GetRegisteredEventsByUser(userId int64) ([]Event, error) {
	query := `
		SELECT e.id, e.name, e.slug, e.city, e.location,
		       e.date, e.organizerName, e.imageUrl,
		       e.description, e.userId
		FROM events e
		INNER JOIN registrations r ON e.id = r.eventId
		WHERE r.userId = ?
			  ORDER BY id DESC
		
	`

	rows, err := db.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event

		err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.Slug,
			&event.City,
			&event.Location,
			&event.Date,
			&event.OrganizerName,
			&event.ImageUrl,
			&event.Description,
			&event.UserId,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}




func SeedEvents( ) error {


	events := []Event{
	{
		Name:          "Summer Music Fest",
		Slug:          "summer-music-fest-2026",
		City:          "Austin",
		Location:      "Zilker Park",
		Date:          time.Date(2026, 8, 15, 18, 0, 0, 0, time.UTC),
		OrganizerName: "Summer Vibes Productions",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "An amazing summer celebration with live bands, food trucks, and outdoor activities. Experience the best of Austin's music scene under the stars.",
		UserId:        1,
	},
	{
		Name:          "Outdoor Art Market",
		Slug:          "outdoor-art-market-seattle",
		City:          "Seattle",
		Location:      "Capitol Hill Park",
		Date:          time.Date(2026, 8, 22, 10, 0, 0, 0, time.UTC),
		OrganizerName: "Local Artists Collective",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Support local artists and discover unique handmade art pieces. A vibrant outdoor market featuring paintings, sculptures, jewelry, and more.",
		UserId:        1,
	},
	{
		Name:          "Summer Cinema Under Stars",
		Slug:          "summer-cinema-austin",
		City:          "Austin",
		Location:      "Barton Springs Pool Grounds",
		Date:          time.Date(2026, 8, 28, 20, 0, 0, 0, time.UTC),
		OrganizerName: "Outdoor Cinema Society",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Enjoy classic films under the night sky. Bring blankets and snacks for a magical outdoor movie experience with friends and family.",
		UserId:        1,
	},
	{
		Name:          "Early Fall Festival",
		Slug:          "early-fall-festival",
		City:          "Seattle",
		Location:      "Green Lake Park",
		Date:          time.Date(2026, 9, 5, 11, 0, 0, 0, time.UTC),
		OrganizerName: "Seasonal Events Inc.",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Celebrate the arrival of fall with live music, pumpkin patches, and seasonal crafts. Perfect for the whole family.",
		UserId:        1,
	},
	{
		Name:          "Tech Talk Tuesday",
		Slug:          "tech-talk-tuesday",
		City:          "Austin",
		Location:      "Impact Hub",
		Date:          time.Date(2026, 9, 17, 19, 0, 0, 0, time.UTC),
		OrganizerName: "Austin Tech Community",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Join us for an evening of tech talks and networking. Industry leaders share insights on the latest trends in software development and AI.",
		UserId:        1,
	},
	{
		Name:          "Yoga and Mindfulness Retreat",
		Slug:          "yoga-mindfulness-retreat",
		City:          "Austin",
		Location:      "Nature's Peace Center",
		Date:          time.Date(2026, 9, 19, 8, 0, 0, 0, time.UTC),
		OrganizerName: "Wellness Mind Body",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Start your day with rejuvenating yoga sessions and mindfulness meditation. Perfect for stress relief and personal well-being.",
		UserId:        1,
	},
	{
		Name:          "Local Brewery Tour",
		Slug:          "local-brewery-tour",
		City:          "Austin",
		Location:      "Austin Craft Brewery District",
		Date:          time.Date(2026, 9, 24, 14, 0, 0, 0, time.UTC),
		OrganizerName: "Beer Lovers United",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Explore Austin's craft brewery scene with guided tours. Taste unique brews, learn brewing techniques, and meet fellow beer enthusiasts.",
		UserId:        1,
	},
	{
		Name:          "Fall Market Fest",
		Slug:          "fall-market-fest",
		City:          "Seattle",
		Location:      "Pike Place Market",
		Date:          time.Date(2026, 9, 26, 10, 0, 0, 0, time.UTC),
		OrganizerName: "Market Collective",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Experience Seattle's iconic market with special fall offerings. Fresh produce, local crafts, and live entertainment throughout the day.",
		UserId:        1,
	},
	{
		Name:          "Startup Networking Night",
		Slug:          "startup-networking-night",
		City:          "Austin",
		Location:      "Domain Innovation Center",
		Date:          time.Date(2026, 9, 29, 18, 30, 0, 0, time.UTC),
		OrganizerName: "Austin Startup Hub",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Connect with founders, investors, and entrepreneurs. A perfect opportunity to pitch ideas, find co-founders, and build your network.",
		UserId:        1,
	},
	{
		Name:          "Virtual Reality Gaming Expo",
		Slug:          "vr-gaming-expo",
		City:          "Austin",
		Location:      "Tech Hub Austin",
		Date:          time.Date(2026, 10, 20, 16, 0, 0, 0, time.UTC),
		OrganizerName: "Gaming Revolution",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Immerse yourself in the future of gaming with cutting-edge VR experiences. Try the latest games, meet developers, and experience next-generation gaming technology.",
		UserId:        1,
	},
	{
		Name:          "Autumn Harvest Fair",
		Slug:          "autumn-harvest-fair",
		City:          "Seattle",
		Location:      "Woodland Park",
		Date:          time.Date(2026, 10, 10, 11, 0, 0, 0, time.UTC),
		OrganizerName: "Harvest Celebrations",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Celebrate the harvest season with pumpkins, fall crafts, live music, and delicious seasonal food. A family-friendly autumn tradition.",
		UserId:        1,
	},
	{
		Name:          "Photography Workshop",
		Slug:          "photography-workshop",
		City:          "Austin",
		Location:      "Downtown Art Studios",
		Date:          time.Date(2026, 10, 15, 14, 0, 0, 0, time.UTC),
		OrganizerName: "Lens Masters Academy",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Learn photography fundamentals from professional photographers. Improve composition, lighting, and editing techniques in this hands-on workshop.",
		UserId:        1,
	},
	{
		Name:          "Podcast Creators Summit",
		Slug:          "podcast-creators-summit",
		City:          "Austin",
		Location:      "Media Production Hub",
		Date:          time.Date(2026, 11, 5, 14, 0, 0, 0, time.UTC),
		OrganizerName: "Podcast Pros Association",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Connect with podcast creators and audio producers. Learn production techniques, monetization strategies, and how to grow your podcast audience.",
		UserId:        1,
	},
	{
		Name:          "Halloween Costume Party",
		Slug:          "halloween-costume-party",
		City:          "Austin",
		Location:      "Downtown Nightclub",
		Date:          time.Date(2026, 10, 31, 20, 0, 0, 0, time.UTC),
		OrganizerName: "Party Central",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Show off your most creative costume at our spooky Halloween party. DJ sets, costume contests, and a night of fun and celebration.",
		UserId:        1,
	},
	{
		Name:          "Winter Holiday Market",
		Slug:          "winter-holiday-market",
		City:          "Seattle",
		Location:      "Seattle Center",
		Date:          time.Date(2026, 11, 20, 10, 0, 0, 0, time.UTC),
		OrganizerName: "Holiday Markets Inc.",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Shop for unique holiday gifts at our festive winter market. Featuring local vendors, holiday decorations, seasonal treats, and live entertainment.",
		UserId:        1,
	},
	{
		Name:          "Nutrition and Health Summit",
		Slug:          "nutrition-health-summit",
		City:          "Austin",
		Location:      "Austin Health Complex",
		Date:          time.Date(2026, 11, 25, 10, 0, 0, 0, time.UTC),
		OrganizerName: "Healthy Lifestyle Institute",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Learn about nutrition science and optimal health strategies. Hear from nutritionists, doctors, and wellness experts on how to live your best life.",
		UserId:        1,
	},
	{
		Name:          "Year End Tech Conference",
		Slug:          "year-end-tech-conference",
		City:          "Austin",
		Location:      "Austin Convention Center",
		Date:          time.Date(2026, 12, 5, 9, 0, 0, 0, time.UTC),
		OrganizerName: "Tech Industry Leaders",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Wrap up the year with keynotes from industry leaders, panel discussions on tech trends, and networking opportunities with innovators worldwide.",
		UserId:        1,
	},
	{
		Name:          "Craft Brewers Championship",
		Slug:          "craft-brewers-championship",
		City:          "Austin",
		Location:      "Austin Brewery District",
		Date:          time.Date(2026, 12, 10, 14, 0, 0, 0, time.UTC),
		OrganizerName: "American Brewers Alliance",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Celebrate the craft beer community at our championship competition. Taste award-winning brews, meet brewmasters, and discover innovative brewing techniques.",
		UserId:        1,
	},
	{
		Name:          "New Year Celebration Gala",
		Slug:          "new-year-celebration-gala",
		City:          "Seattle",
		Location:      "The Paramount Theatre",
		Date:          time.Date(2026, 12, 31, 21, 0, 0, 0, time.UTC),
		OrganizerName: "Celebration Productions",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Celebrate the New Year in style with live music, champagne, and dancing. Ring in 2027 with style at this elegant gala event.",
		UserId:        1,
	},
	{
		Name:          "Plant-Based Food Festival",
		Slug:          "plant-based-food-festival",
		City:          "Seattle",
		Location:      "Magnolia Park",
		Date:          time.Date(2027, 1, 22, 11, 0, 0, 0, time.UTC),
		OrganizerName: "Vegan Collective",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Celebrate plant-based living with delicious food, cooking demonstrations, and nutrition workshops. Discover how to thrive on a vegan diet.",
		UserId:        1,
	},
	{
		Name:          "Machine Learning Bootcamp",
		Slug:          "machine-learning-bootcamp",
		City:          "Austin",
		Location:      "Austin Tech Academy",
		Date:          time.Date(2027, 1, 15, 8, 0, 0, 0, time.UTC),
		OrganizerName: "AI Innovators",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Intensive bootcamp on machine learning and AI. Learn from industry experts, work on real projects, and build a portfolio in cutting-edge AI technologies.",
		UserId:        1,
	},
	{
		Name:          "Mindfulness and Meditation Retreat",
		Slug:          "mindfulness-meditation-retreat",
		City:          "Seattle",
		Location:      "Mountain Vista Wellness Center",
		Date:          time.Date(2027, 2, 20, 7, 0, 0, 0, time.UTC),
		OrganizerName: "Zen Institute",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "A transformative retreat focused on mindfulness practices and meditation. Reconnect with yourself and find inner peace in a serene natural setting.",
		UserId:        1,
	},
	{
		Name:          "Blockchain and Crypto Conference",
		Slug:          "blockchain-crypto-conference",
		City:          "Austin",
		Location:      "Downtown Austin Conference Hall",
		Date:          time.Date(2027, 2, 10, 9, 0, 0, 0, time.UTC),
		OrganizerName: "Crypto Leaders Inc.",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Learn about the latest developments in blockchain and cryptocurrency. Network with industry experts, investors, and innovators in the Web3 space.",
		UserId:        1,
	},
	{
		Name:          "Urban Gardening Workshop",
		Slug:          "urban-gardening-workshop",
		City:          "Seattle",
		Location:      "Community Garden Center",
		Date:          time.Date(2027, 3, 5, 10, 0, 0, 0, time.UTC),
		OrganizerName: "Green City Initiative",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Learn how to grow fresh produce in urban spaces. Tips on container gardening, vertical gardens, and sustainable food production in the city.",
		UserId:        1,
	},
	{
		Name:          "Sustainable Fashion Week",
		Slug:          "sustainable-fashion-week",
		City:          "Seattle",
		Location:      "Seattle Convention Center",
		Date:          time.Date(2027, 3, 15, 10, 0, 0, 0, time.UTC),
		OrganizerName: "EcoStyle Collective",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Explore sustainable and eco-friendly fashion. Discover designers committed to ethical production and environmental conservation. Learn how to build a conscious wardrobe.",
		UserId:        1,
	},
	{
		Name:          "Esports Championship Finals",
		Slug:          "esports-championship-finals",
		City:          "Austin",
		Location:      "Austin Arena Complex",
		Date:          time.Date(2027, 4, 10, 12, 0, 0, 0, time.UTC),
		OrganizerName: "Global Esports League",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Watch the most competitive esports teams battle for championship glory. Experience intense gameplay, legendary moments, and the thrill of professional gaming.",
		UserId:        1,
	},
	{
		Name:          "Women in Tech Conference",
		Slug:          "women-in-tech-conference",
		City:          "Austin",
		Location:      "Tech Center Downtown",
		Date:          time.Date(2027, 4, 22, 9, 0, 0, 0, time.UTC),
		OrganizerName: "Tech Women Empowerment",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Celebrate and empower women in technology. Network with female leaders, engineers, and entrepreneurs. Discuss challenges and opportunities in tech.",
		UserId:        1,
	},
	{
		Name:          "Indie Film Festival",
		Slug:          "indie-film-festival",
		City:          "Seattle",
		Location:      "Seattle Film Society",
		Date:          time.Date(2027, 5, 18, 19, 0, 0, 0, time.UTC),
		OrganizerName: "Independent Cinema Collective",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Discover groundbreaking independent films from emerging filmmakers worldwide. Experience unique stories and connect with the indie film community.",
		UserId:        1,
	},
	{
		Name:          "Marine Conservation Summit",
		Slug:          "marine-conservation-summit",
		City:          "Seattle",
		Location:      "Waterfront Environmental Center",
		Date:          time.Date(2027, 5, 30, 10, 0, 0, 0, time.UTC),
		OrganizerName: "Ocean Defenders Alliance",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Learn about ocean conservation and marine ecosystem protection. Network with environmental scientists and conservation leaders dedicated to saving our oceans.",
		UserId:        1,
	},
	{
		Name:          "Street Art and Graffiti Festival",
		Slug:          "street-art-graffiti-festival",
		City:          "Austin",
		Location:      "Downtown Street Art District",
		Date:          time.Date(2027, 6, 12, 12, 0, 0, 0, time.UTC),
		OrganizerName: "Urban Art Collective",
		ImageUrl:      "https://images.unsplash.com/photo-1625467769506-8742f12fb865?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Witness live street art performances and graffiti masterpieces. Meet talented artists and see the streets transform into a vibrant outdoor gallery.",
		UserId:        1,
	},
	{
		Name:          "Harmony Festival 2027",
		Slug:          "harmony-festival-2027",
		City:          "Austin",
		Location:      "Austin Convention Center",
		Date:          time.Date(2027, 11, 15, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Music Enthusiasts LLC",
		ImageUrl:      "https://images.unsplash.com/photo-1700348305514-b9eea0c9999b?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Harmony Festival is a celebration of all music genres, bringing together musicians, artists, and music enthusiasts from around the world. Experience a day filled with live performances, interactive workshops, and a vibrant atmosphere of creativity and harmony.",
		UserId:        1,
	},
	{
		Name:          "3D Animation Workshop 2027",
		Slug:          "3d-animation-workshop-2027",
		City:          "Austin",
		Location:      "Austin Convention Center",
		Date:          time.Date(2027, 12, 8, 0, 0, 0, 0, time.UTC),
		OrganizerName: "3D Animators Inc.",
		ImageUrl:      "https://images.unsplash.com/photo-1683740128079-09405ab63ab8?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Dive into the captivating world of 3D animation at our exclusive 3D Animation Masterclass! Whether you're an aspiring animator, a student studying animation, or a professional looking to enhance your skills, this workshop offers a unique opportunity to learn from industry experts.",
		UserId:        1,
	},
	{
		Name:          "Rock the City Concert 2027",
		Slug:          "rock-the-city-concert-2027",
		City:          "Austin",
		Location:      "Austin Music Hall",
		Date:          time.Date(2027, 11, 18, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Rock On Productions",
		ImageUrl:      "https://images.unsplash.com/photo-1529787184525-7d3933bec5a0?q=80&w=2831&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1ofHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Get ready to rock out at Rock the City Concert! Experience electrifying performances by top rock bands, enjoy high-energy music, and immerse yourself in an unforgettable night of pure rock and roll.",
		UserId:        1,
	},
	{
		Name:          "Artisan Craft Fair 2027",
		Slug:          "artisan-craft-fair-2027",
		City:          "Seattle",
		Location:      "Seattle Exhibition Center",
		Date:          time.Date(2027, 12, 1, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Craftsmanship Guild",
		ImageUrl:      "https://images.unsplash.com/photo-1642178225043-f299072af862?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=2070&q=100",
		Description:   "Discover unique handmade crafts and artworks at the Artisan Craft Fair. Meet talented artisans, shop for one-of-a-kind items, and support local craftsmanship.",
		UserId:        1,
	},
	{
		Name:          "Jazz Fusion Night 2027",
		Slug:          "jazz-fusion-night-2027",
		City:          "Austin",
		Location:      "Austin Jazz Lounge",
		Date:          time.Date(2027, 11, 29, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Groove Masters Productions",
		ImageUrl:      "https://images.unsplash.com/photo-1415201364774-f6f0bb35f28f?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Indulge in the smooth melodies and rhythmic beats of jazz fusion at Jazz Fusion Night. Experience world-class jazz performances, savor delicious cocktails, and immerse yourself in the soulful ambiance of live jazz music.",
		UserId:        1,
	},
	{
		Name:          "Indie Music Showcase 2027",
		Slug:          "indie-music-showcase-2027",
		City:          "Austin",
		Location:      "Austin Indie Spot",
		Date:          time.Date(2027, 11, 25, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Indie Vibes Records",
		ImageUrl:      "https://images.unsplash.com/photo-1470225620780-dba8ba36b745?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Discover the next big indie artists at the Indie Music Showcase. Experience live performances by emerging talents, support independent music, and be part of a vibrant community of music enthusiasts and artists.",
		UserId:        1,
	},
	{
		Name:          "Global Food Festival 2027",
		Slug:          "global-food-festival-2027",
		City:          "Seattle",
		Location:      "Seattle Waterfront Park",
		Date:          time.Date(2027, 10, 30, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Foodie Ventures Inc.",
		ImageUrl:      "https://images.unsplash.com/photo-1504754524776-8f4f37790ca0?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Embark on a culinary journey around the world at the Global Food Festival. Delight your taste buds with international cuisines, cooking demonstrations, and food tastings. Experience the flavors of different cultures in one delicious event.",
		UserId:        1,
	},
	{
		Name:          "Tech Innovators Summit 2027",
		Slug:          "tech-innovators-summit-2027",
		City:          "Seattle",
		Location:      "Seattle Convention Center",
		Date:          time.Date(2027, 11, 15, 0, 0, 0, 0, time.UTC),
		OrganizerName: "InnovateTech Inc.",
		ImageUrl:      "https://images.unsplash.com/photo-1486312338219-ce68d2c6f44d?q=80&w=2944&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fhx8fGVufDB8fHx8fA%3D%3D",
		Description:   "The Tech Innovators Summit is where visionaries, entrepreneurs, and tech enthusiasts converge. Explore the latest technological advancements, attend insightful keynotes from industry leaders, and participate in hands-on workshops.",
		UserId:        1,
	},
	{
		Name:          "Enchanted Garden Gala 2027",
		Slug:          "enchanted-garden-gala-2027",
		City:          "Austin",
		Location:      "Austin Museum of Art",
		Date:          time.Date(2027, 12, 2, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Cultural Garden Society",
		ImageUrl:      "https://images.unsplash.com/photo-1460533893735-45cea2212645?q=80&w=3028&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fhx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Step into a world of wonder at the Enchanted Garden Gala, a magical evening of art, music, and fantasy. Explore enchanting garden installations, experience live performances by world-class musicians and dancers.",
		UserId:        1,
	},
	{
		Name:          "Comedy Extravaganza 2027",
		Slug:          "comedy-extravaganza-2027",
		City:          "Austin",
		Location:      "Austin Laugh Factory",
		Date:          time.Date(2027, 11, 6, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Laugh Productions",
		ImageUrl:      "https://images.unsplash.com/photo-1683117855296-979f17e62e87?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fhx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Prepare for a night of laughter with top comedians from around the world. Enjoy stand-up, improv, and sketches that will have you in stitches!",
		UserId:        1,
	},
	{
		Name:          "Science and Space Expo 2027",
		Slug:          "science-space-expo-2027",
		City:          "Seattle",
		Location:      "Seattle Science Center",
		Date:          time.Date(2027, 10, 29, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Cosmic Explorers Society",
		ImageUrl:      "https://images.unsplash.com/photo-1493528237448-144452699e16?q=80&w=2803&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fhx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Explore the wonders of science and space at this interactive expo. Engage in hands-on experiments, meet scientists, and learn about the mysteries of the universe.",
		UserId:        1,
	},
	{
		Name:          "Fashion Runway 2027",
		Slug:          "fashion-runway-2027",
		City:          "Austin",
		Location:      "Austin Fashion Week Venue",
		Date:          time.Date(2027, 11, 12, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Chic Trends Agency",
		ImageUrl:      "https://images.unsplash.com/photo-1708660369108-2159953bf436?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fhx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Witness the latest trends on the runway. Top designers will showcase their collections, setting the stage for the future of fashion.",
		UserId:        1,
	},
	{
		Name:          "Culinary Masterclass 2027",
		Slug:          "culinary-masterclass-2027",
		City:          "Seattle",
		Location:      "Seattle Epicurean Institute",
		Date:          time.Date(2027, 12, 2, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Gourmet Chefs Society",
		ImageUrl:      "https://images.unsplash.com/photo-1625937712159-e305336cbf4b?q=80&w=2389&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8fHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Join renowned chefs for a culinary journey. Learn cooking techniques, taste exquisite dishes, and elevate your skills in the art of gastronomy.",
		UserId:        1,
	},
	{
		Name:          "Film Buffs Symposium 2027",
		Slug:          "film-buffs-symposium-2027",
		City:          "Austin",
		Location:      "Austin Film Institute",
		Date:          time.Date(2027, 11, 8, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Cinema Society",
		ImageUrl:      "https://images.unsplash.com/photo-1518929458119-e5bf444c30f4?q=80&w=2874&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Description:   "A gathering for film enthusiasts! Screen classic movies, engage in discussions with filmmakers, and gain insights into the world of cinema.",
		UserId:        1,
	},
	{
		Name:          "Literary Salon 2027",
		Slug:          "literary-salon-2027",
		City:          "Seattle",
		Location:      "Seattle & Co. Bookstore",
		Date:          time.Date(2027, 12, 15, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Words Society",
		ImageUrl:      "https://images.unsplash.com/photo-1518373714866-3f1478910cc0?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fhx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Celebrate the written word at this literary gathering. Listen to readings by acclaimed authors, participate in book discussions, and embrace the magic of storytelling.",
		UserId:        1,
	},
	{
		Name:          "Wellness Expo 2027",
		Slug:          "wellness-expo-2027",
		City:          "Austin",
		Location:      "Austin Convention Center",
		Date:          time.Date(2027, 11, 30, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Wellness Warriors Inc.",
		ImageUrl:      "https://images.unsplash.com/photo-1545205597-3d9d02c29597?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fhx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Immerse yourself in the world of fitness and well-being. Attend fitness classes, learn about nutrition, and explore holistic approaches to health.",
		UserId:        1,
	},
	{
		Name:          "Digital Art Symposium 2027",
		Slug:          "digital-art-symposium-2027",
		City:          "Seattle",
		Location:      "Seattle Art Gallery",
		Date:          time.Date(2027, 11, 1, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Tech Creatives Collective",
		ImageUrl:      "https://images.unsplash.com/photo-1459908676235-d5f02a50184b?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fhx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Discover the intersection of technology and art. Experience digital art installations, attend VR workshops, and meet digital artists pushing creative boundaries.",
		UserId:        1,
	},
	{
		Name:          "Dance Fusion Festival 2027",
		Slug:          "dance-fusion-festival-2027",
		City:          "Austin",
		Location:      "Austin Street Dance Studio",
		Date:          time.Date(2027, 11, 28, 0, 0, 0, 0, time.UTC),
		OrganizerName: "Rhythm Revolution",
		ImageUrl:      "https://images.unsplash.com/photo-1504609813442-a8924e83f76e?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fhx8fGVufDB8fHx8fA%3D%3D",
		Description:   "Experience a blend of dance styles from around the world. Participate in dance workshops, watch electrifying performances, and dance the night away.",
		UserId:        1,
	},
}

	query := `
		INSERT INTO events (
			name,
			slug,
			city,
			location,
			date,
			organizerName,
			imageUrl,
			description,
			userId
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	for _, event := range events {
		result, err := db.DB.Exec(
			query,
			event.Name,
			event.Slug,
			event.City,
			event.Location,
			event.Date,
			event.OrganizerName,
			event.ImageUrl,
			event.Description,
			event.UserId,
		)

		if err != nil {
			return fmt.Errorf("failed to seed event %q: %w", event.Name, err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			return fmt.Errorf("failed to get inserted event id: %w", err)
		}

		fmt.Printf("created event with id: %d (userId: %d)\n", id, event.UserId)
	}

	fmt.Println("seeding finished.")

	return nil
}
