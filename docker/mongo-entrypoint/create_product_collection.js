db = db.getSiblingDB('poc_hexagonal_db');

db.createCollection('products')

db.products.insertMany([
    { name: 'IPHONE' },
    { name: 'LAPTOP I9' },
    { name: 'TV SAMSUNG' }
]);