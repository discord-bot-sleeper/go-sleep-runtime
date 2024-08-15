curl requ:

curl -X POST http://localhost:3000/add \
-H "Content-Type: application/json" \
-d '{"token":"MTI2MjgzMDEzNTI4Mzk0MTM3Ng.G7Aeab.h8UNc_Z34vRARjI-l2bIAXc_wAdvZMntSzS1-w","uuid":"asdasd"}'

bot2
curl -X POST http://localhost:3333/add -H "Content-Type: application/json" -d '{"token": "MTI2MjgzMDM0NzcwNDQ2NzUxNw.Glybex.HX4iYAIFr4_29sTvI3hSBQF6aul1zVtr23IA00", "uuid": "1234"}'

bot1
curl -X POST http://localhost:3333/add -H "Content-Type: application/json" -d '{"token": "MTI2MjgzMDEzNTI4Mzk0MTM3Ng.G7Aeab.h8UNc_Z34vRARjI-l2bIAXc_wAdvZMntSzS1-w", "uuid": "bot1"}'
