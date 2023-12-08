workspace "Online Shopping Platform" "An online shopping platform to facilitate e-commerce transactions for various products." {

    model {
        customer = person "Customer" "A user who browses and purchases products."
        vendor = person "Vendor" "A supplier who provides products to sell on the platform."
        admin = person "Administrator" "Operates and manages the online shopping platform."

        onlineShoppingPlatform = softwaresystem "Online Shopping Platform" "Facilitates e-commerce transactions for various products." {
            webApplication = container "Web Application" "Allows customers to browse and purchase products." "React/Angular"
            mobileApp = container "Mobile App" "A mobile application for shopping on the go." "Flutter/React Native"
            adminPortal = container "Admin Portal" "Web-based interface for platform management." "Angular/React"
            apiGateway = container "API Gateway" "Handles all incoming requests and routes them to appropriate services." "Spring Cloud Gateway"
            
            serviceLayer = container "Service Layer" "Contains business logic for the platform." {
                userAuthService = component "User Authentication Service" "Manages user sign-in and security." "Spring Boot"
                productCatalogueService = component "Product Catalogue Service" "Handles product listings and categories." "Spring Boot"
                orderManagementService = component "Order Management Service" "Processes orders and payments." "Spring Boot"
                inventoryManagementService = component "Inventory Management Service" "For vendors to manage their inventory." "Spring Boot"
                paymentProcessingService = component "Payment Processing Service" "Handles payment transactions." "Spring Boot"
                reviewRatingService = component "Review and Rating Service" "Manages customer reviews and ratings." "Spring Boot"
            }

            database = container "Database" "Stores user data, product information, orders, etc." "SQL/NoSQL Database"
        }

        // relationships
        customer -> webApplication "Uses"
        customer -> mobileApp "Uses"
        vendor -> adminPortal "Uses"
        admin -> adminPortal "Uses"
        webApplication -> apiGateway "Interacts with"
        mobileApp -> apiGateway "Interacts with"
        apiGateway -> serviceLayer "Routes requests to"
        serviceLayer -> database "Reads from and writes to"
    }

    views {
        systemContext onlineShoppingPlatform "SystemContext" {
            include *
            autoLayout
        }

        container onlineShoppingPlatform "Containers" {
            include *
            autoLayout
        }

        component serviceLayer "Components" {
            include *
            autoLayout
        }

        styles {
            element "Software System" {
                background #1168bd
                color #ffffff
            }
            element "Person" {
                shape person
                background #08427b
                color #ffffff
            }
            element "Container" {
                background #0e72f4
                color #ffffff
            }
            element "Component" {
                background #66adf4
                color #ffffff
            }
            element "Database" {
                shape Cylinder
            }
        }
    }
}
