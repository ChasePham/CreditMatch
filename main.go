package main


import ("fmt"
"CreditMatch/packages/amex"
"CreditMatch/packages/chase"
"CreditMatch/packages/capital_one"
"CreditMatch/packages/citi"
)


func main() {
    fmt.Println("Welcome to Credit Match! \nInput your statement information to figure out which credit card bundle works for you!")


    fmt.Print("How much do you spend a year in spending categories with Citi Bank: ")
    var citi_spending_cat float64
    fmt.Scanln(&citi_spending_cat)

    fmt.Print("How much do you spend a year in seasonal categories within Chase Bank: ")
    var chase_seasonal_cat float64
    fmt.Scanln(&chase_seasonal_cat)

    fmt.Print("How much do you spend a year in air travel through a Bank Portal: ")
    var air_travel_portal float64
    fmt.Scanln(&air_travel_portal)

    fmt.Print("How much do you spend a year in other air travel purchases: ")
    var air_other float64
    fmt.Scanln(&air_other)

    fmt.Print("How much do you spend a year in hotels and rentals through a Bank Portal: ")
    var hotels_and_rentals_portal float64
    fmt.Scanln(&hotels_and_rentals_portal)

    fmt.Print("How much do you spend a year in other hotels and rentals: ")
    var hotels_and_rentals_other float64
    fmt.Scanln(&hotels_and_rentals_other)

    fmt.Print("How much do you spend a year at gas stations: ")
    var gas float64
    fmt.Scanln(&gas)


    fmt.Print("How much do you spend a year in ubers: ")
    var ubers float64
    fmt.Scanln(&ubers)


    fmt.Print("How much do you spend a year at drug stores: ")
    var drugstores float64
    fmt.Scanln(&drugstores)

    fmt.Print("How much do you spend a year at restaurants: ")
    var restaurants float64
    fmt.Scanln(&restaurants)

    fmt.Print("How much do you spend a year in online groceries: ")
    var online_groceries float64
    fmt.Scanln(&online_groceries)

    fmt.Print("How much do you spend a year in groceries: ")
    var groceries float64
    fmt.Scanln(&groceries)

    fmt.Print("How much do you spend a year in streaming services: ")
    var streaming_services float64
    fmt.Scanln(&streaming_services)

    fmt.Print("How much do you spend a year in entertainment: ")
    var entertainment float64
    fmt.Scanln(&entertainment)

    fmt.Print("How much do you spend a year in retail: ")
    var retail float64
    fmt.Scanln(&retail)

    fmt.Print("How much do you spend a year in other purchases: ")
    var other_purchases float64
    fmt.Scanln(&other_purchases)

    // FIXME: Forgot other purchases
    amex_trifecta := amex.Amex {
        Supermarkets: online_groceries + online_groceries,
        Online_retail: retail,
        Gas_stations: gas,
        Restaurants: restaurants,
        Flights_portal: air_travel_portal,
        Hotels: hotels_and_rentals_other + hotels_and_rentals_portal,
        Ubers: ubers,
        Digital_entertainment: streaming_services,
    }

    chase_trifecta := chase.Chase {
        Flex_cat : chase_seasonal_cat,
        Travel_portal_purchases :hotels_and_rentals_portal + air_travel_portal,
        Other_travel_purchases : hotels_and_rentals_other + air_other,
        Drug_stores : drugstores,
        Restaurants :restaurants,
        Online_groceries : online_groceries,
        Streaming_services : streaming_services,
        Other_purchases : other_purchases,
        Hotels_and_car_rentals : hotels_and_rentals_other + hotels_and_rentals_portal,
    }

    capital_one_duo := capital_one.Capital {
        Dining: restaurants,
        Entertainment: entertainment,
        Streaming_services: streaming_services,
        Groceries: online_groceries + groceries,
        Hotels_and_car_rentals: hotels_and_rentals_other + hotels_and_rentals_portal,
        Flights: air_travel_portal + air_other,
    }

    citi_trifecta := citi.Citi {
        Other_purchases : other_purchases,
        Custom_categories : citi_spending_cat,
        Travel_portal_purchases : hotels_and_rentals_portal,
        Other_travel_purchases : hotels_and_rentals_other + air_other,
    }
}
