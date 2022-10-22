// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

"use strict";

const express = require("express");
const fs = require("fs");

const nsConfigRaw = fs.readFileSync("/namespace/config/runtime.json");
const nsConfig = JSON.parse(nsConfigRaw);
console.log(`Namespace config: ${JSON.stringify(nsConfig, null, 2)}`);

// Constants
const PORT = nsConfig.current.port.find((s) => s.name === "webapi").port;
const HOST = "0.0.0.0";

// App
const app = express();
app.get("/", (req, res) => {
	res.send(`Hello, world!`);
});

app.listen(PORT, HOST);

console.log(`Running on http://${HOST}:${PORT}`);
