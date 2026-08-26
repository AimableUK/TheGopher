// scripts/init_indexes.js

// Connect to the database
db = db.getSiblingDB("library");

// Add root user for the MongoDB instance
db.createUser({
  user: "admin",
  pwd: "password",
  roles: [{ role: "root", db: "admin" }],
});

// Schema validation for the books collection
db.createCollection("books", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["title", "author_id", "published_date", "details", "category"],
      properties: {
        title: {
          bsonType: "string",
          description: "Title of the book",
        },
        author_id: {
          bsonType: "objectId",
          description: "Reference to the author",
        },
        published_date: {
          bsonType: "date",
          description: "Publication date",
        },
        details: {
          bsonType: "object",
          description: "Additional details of the book",
        },
        category: {
          bsonType: "string",
          description: "Category of the book",
        },
      },
    },
  },
});

// This creates a text index on the title field and the details.summary field in the books collection.
// enables full-text search on both fields, allowing you to efficiently search for books based on keywords in either the title or summary.
db.books.createIndex({ title: "text", "details.summary": "text" });

// Creates an ascending index on the name field in the authors collection.
// speeds up queries that filter or sort authors by their name, making lookups for specific author names faster.
db.authors.createIndex({ name: 1 });

// Enable sharding on the "library" database
// Activates sharding for the library database, allowing its collections to be distributed across multiple shards for load balancing and scalability.
sh.enableSharding("library");

// Enables sharding specifically for the books collection in the library database.
// Uses the _id field as the shard key, with a hashed distribution. Hashed sharding evenly distributes documents across shards, which is beneficial for write-heavy workloads by minimizing "hot spots" on specific shards.
sh.shardCollection("library.books", { _id: "hashed" });
