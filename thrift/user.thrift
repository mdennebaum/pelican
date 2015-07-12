namespace go user

struct User {
  1: required string id,
  2: required string name,
  3: required string screenname,
  4: required string bio,
  5: required string phone,
  6: optional string email,
  7: required string created,
  8: required string updated
}

service UserSvc {
  User create(1:User user),
  User read(1:string userId),
  User update(1:User user),
  void destroy(1:string userId),
  list<User> fetch(),
  void reset()
}