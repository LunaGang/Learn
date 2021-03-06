package main

type Person struct {
    name string
    age int
    city,phone string
}

type Speaker struct {
    Person //type embedding for composition
    speaksOn []string
    pastEvents []string
}
 
type Organizer struct {
    Person //type embedding for composition
    meetups []string
}
 
type Attendee struct {
    Person //type embedding for composition
    interests []string
}

/**
 * Speaker reçois les propriétés/méthodes de Person, Organizer et Attendee aussi. C'est une sorte d'héritage, les valeurs de Person sont appelé dedans.
 * 
 * Et si Person et Speaker on des propriétés en commun, vous pouvez faire :
 * speaker.Person.Method()
 *
 * Au lieu de :
 * speaker.Method(), si ils en avaient pas.
**/
