const MAX_PERSONS = process.env.MOCK_PERSONS ? parseInt(process.env.MOCK_PERSONS, 10) : 100;
const DB_NAME = process.env.MONGO_INITDB_DATABASE;
const COLLECTION_NAME = process.env.COLLECTION_NAME;

print("Start of init-mongo.js script");
print(`Creating '${COLLECTION_NAME}' collection in database '${DB_NAME}' with ${MAX_PERSONS} fake persons.`);

db = db.getSiblingDB(DB_NAME);

db.createCollection(COLLECTION_NAME, {
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

for (let i = 1; i <= MAX_PERSONS; i++) {
  const name = names[Math.floor(Math.random() * names.length)];
  persons.push({
    id: i,
    name: name,
    age: Math.floor(Math.random() * 100),
    email: `person${i}@example.com`
  });
}

db[COLLECTION_NAME].insertMany(persons);

print(`Inserted ${persons.length} persons into the '${COLLECTION_NAME}' collection.`);
print("End of init-mongo.js script");