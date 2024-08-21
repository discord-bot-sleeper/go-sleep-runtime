ADD metodit

bot1
curl -X POST http://localhost:8080/add -H "Content-Type: application/json" -d '{"token": "MTI2MjgzMDEzNTI4Mzk0MTM3Ng.G7Aeab.h8UNc_Z34vRARjI-l2bIAXc_wAdvZMntSzS1-w", "uuid": "bot1"}'

bot2
curl -X POST http://localhost:8080/add -H "Content-Type: application/json" -d '{"token": "MTI2MjgzMDM0NzcwNDQ2NzUxNw.Glybex.HX4iYAIFr4_29sTvI3hSBQF6aul1zVtr23IA00", "uuid": "1234"}'

REMOVE metodit

bot1
curl -X POST http://localhost:8080/remove -H "Content-Type: application/json" -d '{"uuid": "bot1"}'

bot2
curl -X POST http://localhost:8080/remove -H "Content-Type: application/json" -d '{"uuid": "1234"}'

CURRENT metodi: curl http://localhost:8080/current
LIST metodi: curl http://localhost:8080/list
