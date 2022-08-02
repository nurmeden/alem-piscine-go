student@ALEM-316f41:~$ curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"Dulat_Nurmeden\"}}){id}}"}'
{"data":{"user":[{"id":9619}]}}student@ALEM-316f41:~$ curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"q
{lat_Nurmeden\"}}){id}}"}' | jq 
  "data": {
    "user": [
      {
        "id": 9619
      }
    ]
  }
}
student@ALEM-316f41:~$ curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"Dulat_Nurmeden\"}}){id}}"}' | jq".data"\
> exit
jq.dataexit: command not found
student@ALEM-316f41:~$ curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"Dulat_Nurmeden\"}}){id}}"}' | jq".data"
jq.data: command not found
student@ALEM-316f41:~$ curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"Dulat_Nurmeden\"}}){id}}"}' | jq '.data'
{
  "user": [
    {
      "id": 9619
    }
  ]
}
student@ALEM-316f41:~$ curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"Dulat_Nurmeden\"}}){id}}"}' | jq '.data.user[]'
{
  "id": 9619
}
student@ALEM-316f41:~$ curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"Dulat_Nurmeden\"}}){id}}"}' | jq '.data.user[].id'
9619

