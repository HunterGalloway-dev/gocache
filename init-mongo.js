print("Start of init-mongo.js script");

print(process.env.MONGO_INITDB_DATABASE);
print(process.env.COLLECTION_NAME)

db = db.getSiblingDB(process.env.MONGO_INITDB_DATABASE);

db.createCollection('person', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['name', 'age', 'email'],
      properties: {
        name: {
          bsonType: 'string',
          description: 'must be a string and is required'
        },
        age: {
          bsonType: 'int',
          minimum: 0,
          description: 'must be an integer and is required'
        },
        email: {
          bsonType: 'string',
          pattern: '^.+@.+$',
          description: 'must be a string and match the regular expression pattern'
        }
      }
    }
  }
});

const names = [
  "John Doe", "Jane Smith", "Alice Johnson", "Bob Brown", "Charlie Davis",
  "Diana Evans", "Eve Foster", "Frank Green", "Grace Harris", "Hank Irving",
  "Ivy Johnson", "Jack King", "Kara Lee", "Leo Miller", "Mona Nelson",
  "Nina Owens", "Oscar Perry", "Paula Quinn", "Quincy Roberts", "Rita Scott"
];

const persons = [];
for (let i = 1; i <= 10; i++) {
  const name = names[Math.floor(Math.random() * names.length)];
  persons.push({
    name: name,
    age: Math.floor(Math.random() * 100),
    email: `person${i}@example.com`
  });
}

db.person.insertMany(persons);

print("End of init-mongo.js script");