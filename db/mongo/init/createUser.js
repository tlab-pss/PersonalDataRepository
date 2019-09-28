db.createUser(
  {
    user: 'user',
    pwd: 'password',
    roles: [
      {
        role: 'readWrite',
        db: 'pss'
      }
    ]
  }
);

// run during an initialization period and only if the database is empty
