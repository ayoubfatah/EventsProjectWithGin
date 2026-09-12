package models

import (
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

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)
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

	return events, nil
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
	query := `SELECT * FROM events WHERE city = ?`

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


//  seeds

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
 





// func SeedEvents() error {


// 	events := []Event{
// 		{
// 			Name:          "Harmony Festival",
// 			Slug:          "harmony-festival",
// 			City:          "Austin",
// 			Location:      "Austin Convention Center",
// 			Date:          time.Date(2030, 11, 15, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Music Enthusiasts LLC",
// 			ImageUrl:      "https://images.unsplash.com/photo-1700348305514-b9eea0c9999b?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "Harmony Festival is a celebration of all music genres, bringing together musicians, artists, and music enthusiasts from around the world. Experience a day filled with live performances, interactive workshops, and a vibrant atmosphere of creativity and harmony. Join us for an unforgettable musical journey!",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "3D Animation Workshop",
// 			Slug:          "3d-animation-workshop",
// 			City:          "Austin",
// 			Location:      "Austin Convention Center",
// 			Date:          time.Date(2030, 12, 8, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "3D Animators Inc.",
// 			ImageUrl:      "https://images.unsplash.com/photo-1683740128079-09405ab63ab8?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "Dive into the captivating world of 3D animation at our exclusive 3D Animation Masterclass! Whether you're an aspiring animator, a student studying animation, or a professional looking to enhance your skills, this workshop offers a unique opportunity to learn from industry experts and elevate your animation prowess.",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Rock the City Concert",
// 			Slug:          "rock-the-city-concert",
// 			City:          "Austin",
// 			Location:      "Austin Music Hall",
// 			Date:          time.Date(2030, 11, 18, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Rock On Productions",
// 			ImageUrl:      "https://images.unsplash.com/photo-1529787184525-7d3933bec5a0?q=80&w=2831&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1ofHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "Get ready to rock out at Rock the City Concert! Experience electrifying performances by top rock bands, enjoy high-energy music, and immerse yourself in an unforgettable night of pure rock and roll.",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Artisan Craft Fair",
// 			Slug:          "artisan-craft-fair",
// 			City:          "Seattle",
// 			Location:      "Seattle Exhibition Center",
// 			Date:          time.Date(2030, 12, 1, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Craftsmanship Guild",
// 			ImageUrl:      "https://images.unsplash.com/photo-1642178225043-f299072af862?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=2070&q=100",
// 			Description:   "Discover unique handmade crafts and artworks at the Artisan Craft Fair. Meet talented artisans, shop for one-of-a-kind items, and support local craftsmanship. Join us for a day of creativity and craftsmanship.",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Jazz Fusion Night",
// 			Slug:          "jazz-fusion-night",
// 			City:          "Austin",
// 			Location:      "Austin Jazz Lounge",
// 			Date:          time.Date(2030, 11, 29, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Groove Masters Productions",
// 			ImageUrl:      "https://images.unsplash.com/photo-1415201364774-f6f0bb35f28f?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "Indulge in the smooth melodies and rhythmic beats of jazz fusion at Jazz Fusion Night. Experience world-class jazz performances, savor delicious cocktails, and immerse yourself in the soulful ambiance of live jazz music.",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Indie Music Showcase",
// 			Slug:          "indie-music-showcase",
// 			City:          "Austin",
// 			Location:      "Austin Indie Spot",
// 			Date:          time.Date(2030, 11, 25, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Indie Vibes Records",
// 			ImageUrl:      "https://images.unsplash.com/photo-1470225620780-dba8ba36b745?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "Discover the next big indie artists at the Indie Music Showcase. Experience live performances by emerging talents, support independent music, and be part of a vibrant community of music enthusiasts and artists.",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Global Food Festival",
// 			Slug:          "global-food-festival",
// 			City:          "Seattle",
// 			Location:      "Seattle Waterfront Park",
// 			Date:          time.Date(2030, 10, 30, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Foodie Ventures Inc.",
// 			ImageUrl:      "https://images.unsplash.com/photo-1504754524776-8f4f37790ca0?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "Embark on a culinary journey around the world at the Global Food Festival. Delight your taste buds with international cuisines, cooking demonstrations, and food tastings. Experience the flavors of different cultures in one delicious event.",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Tech Innovators Summit",
// 			Slug:          "tech-innovators-summit",
// 			City:          "Seattle",
// 			Location:      "Seattle Convention Center",
// 			Date:          time.Date(2030, 11, 15, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "InnovateTech Inc.",
// 			ImageUrl:      "https://images.unsplash.com/photo-1486312338219-ce68d2c6f44d?q=80&w=2944&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "The Tech Innovators Summit is where visionaries, entrepreneurs, and tech enthusiasts converge. Explore the latest technological advancements, attend insightful keynotes from industry leaders, and participate in hands-on workshops. Connect with innovators, pitch your ideas, and be a part of shaping the future of technology.",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Enchanted Garden Gala",
// 			Slug:          "enchanted-garden-gala",
// 			City:          "Austin",
// 			Location:      "Austin Museum of Art",
// 			Date:          time.Date(2030, 12, 2, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Cultural Garden Society",
// 			ImageUrl:      "https://images.unsplash.com/photo-1460533893735-45cea2212645?q=80&w=3028&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "Step into a world of wonder at the Enchanted Garden Gala, a magical evening of art, music, and fantasy. Explore enchanting garden installations, experience live performances by world-class musicians and dancers, and indulge in gourmet delicacies. Dress in your most glamorous attire and immerse yourself in a night of elegance and enchantment.",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Comedy Extravaganza",
// 			Slug:          "comedy-extravaganza",
// 			City:          "Austin",
// 			Location:      "Austin Laugh Factory",
// 			Date:          time.Date(2030, 11, 6, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Laugh Productions",
// 			ImageUrl:      "https://images.unsplash.com/photo-1683117855296-979f17e62e87?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "Prepare for a night of laughter with top comedians from around the world. Enjoy stand-up, improv, and sketches that will have you in stitches!",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Science and Space Expo",
// 			Slug:          "science-space-expo",
// 			City:          "Seattle",
// 			Location:      "Seattle Science Center",
// 			Date:          time.Date(2030, 10, 29, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Cosmic Explorers Society",
// 			ImageUrl:      "https://images.unsplash.com/photo-1493528237448-144452699e16?q=80&w=2803&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "Explore the wonders of science and space at this interactive expo. Engage in hands-on experiments, meet scientists, and learn about the mysteries of the universe.",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Fashion Runway",
// 			Slug:          "fashion-runway",
// 			City:          "Austin",
// 			Location:      "Austin Fashion Week Venue",
// 			Date:          time.Date(2030, 11, 12, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Chic Trends Agency",
// 			ImageUrl:      "https://images.unsplash.com/photo-1708660369108-2159953bf436?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "Witness the latest trends on the runway. Top designers will showcase their collections, setting the stage for the future of fashion.",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Culinary Masterclass",
// 			Slug:          "culinary-masterclass",
// 			City:          "Seattle",
// 			Location:      "Seattle Epicurean Institute",
// 			Date:          time.Date(2030, 12, 2, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Gourmet Chefs Society",
// 			ImageUrl:      "https://images.unsplash.com/photo-1625937712159-e305336cbf4b?q=80&w=2389&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8fHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "Join renowned chefs for a culinary journey. Learn cooking techniques, taste exquisite dishes, and elevate your skills in the art of gastronomy.",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Film Buffs Symposium",
// 			Slug:          "film-buffs-symposium",
// 			City:          "Austin",
// 			Location:      "Austin Film Institute",
// 			Date:          time.Date(2030, 11, 8, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Cinema Society",
// 			ImageUrl:      "https://images.unsplash.com/photo-1518929458119-e5bf444c30f4?q=80&w=2874&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8fHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "A gathering for film enthusiasts! Screen classic movies, engage in discussions with filmmakers, and gain insights into the world of cinema.",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Literary Salon",
// 			Slug:          "literary-salon",
// 			City:          "Seattle",
// 			Location:      "Seattle & Co. Bookstore",
// 			Date:          time.Date(2030, 12, 15, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Words Society",
// 			ImageUrl:      "https://images.unsplash.com/photo-1518373714866-3f1478910cc0?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "Celebrate the written word at this literary gathering. Listen to readings by acclaimed authors, participate in book discussions, and embrace the magic of storytelling.",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Wellness Expo",
// 			Slug:          "wellness-expo",
// 			City:          "Austin",
// 			Location:      "Austin Convention Center",
// 			Date:          time.Date(2030, 11, 30, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Wellness Warriors Inc.",
// 			ImageUrl:      "https://images.unsplash.com/photo-1545205597-3d9d02c29597?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "Immerse yourself in the world of fitness and well-being. Attend fitness classes, learn about nutrition, and explore holistic approaches to health.",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Digital Art Symposium",
// 			Slug:          "digital-art-symposium",
// 			City:          "Seattle",
// 			Location:      "Seattle Art Gallery",
// 			Date:          time.Date(2030, 11, 1, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Tech Creatives Collective",
// 			ImageUrl:      "https://images.unsplash.com/photo-1459908676235-d5f02a50184b?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "Discover the intersection of technology and art. Experience digital art installations, attend VR workshops, and meet digital artists pushing creative boundaries.",
// 			UserId:        1,
// 		},
// 		{
// 			Name:          "Dance Fusion Festival",
// 			Slug:          "dance-fusion-festival",
// 			City:          "Austin",
// 			Location:      "Austin Street Dance Studio",
// 			Date:          time.Date(2030, 11, 28, 0, 0, 0, 0, time.UTC),
// 			OrganizerName: "Rhythm Revolution",
// 			ImageUrl:      "https://images.unsplash.com/photo-1504609813442-a8924e83f76e?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fHx8fGVufDB8fHx8fA%3D%3D",
// 			Description:   "Experience a blend of dance styles from around the world. Participate in dance workshops, watch electrifying performances, and dance the night away.",
// 			UserId:        1,
// 		},
// 	}

// 	query := `
// 		INSERT INTO events (
// 			name,
// 			slug,
// 			city,
// 			location,
// 			date,
// 			organizerName,
// 			imageUrl,
// 			description,
// 			userId
// 		)
// 		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
// 	`

// 	for _, event := range events {
// 		result, err := db.DB.Exec(
// 			query,
// 			event.Name,
// 			event.Slug,
// 			event.City,
// 			event.Location,
// 			event.Date,
// 			event.OrganizerName,
// 			event.ImageUrl,
// 			event.Description,
// 			event.UserId,
// 		)

// 		if err != nil {
// 			return fmt.Errorf("failed to seed event %q: %w", event.Name, err)
// 		}

// 		id, err := result.LastInsertId()
// 		if err != nil {
// 			return fmt.Errorf("failed to get inserted event id: %w", err)
// 		}

// 		fmt.Printf("created event with id: %d (userId: %d)\n", id, event.UserId)
// 	}

// 	fmt.Println("seeding finished.")

// 	return nil
// }