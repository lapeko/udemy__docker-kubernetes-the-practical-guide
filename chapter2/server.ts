import express from 'express';
import bodyParser from 'body-parser';

const app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

let goal = "";

app.get("/", (req, res) => {
  res.send(`
    <!doctype html>
    <html lang="en">
    <head>
      <meta charset="UTF-8">
      <meta name="viewport"
            content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
      <meta http-equiv="X-UA-Compatible" content="ie=edge">
      <title>Document</title>
    </head>
      <body>
        ${goal ? `<p>${goal}</p>` : ""}
        <form method="POST">
          <input type="text" name="goal">
          <button>Set goal</button>
        </form>
      </body>
    </html>
  `);
});

app.post("/", (req, res) => {
  goal = req.body.goal;
  res.redirect("/");
});

app.listen(3000, () => console.log("Server is running on PORT 3000"));
