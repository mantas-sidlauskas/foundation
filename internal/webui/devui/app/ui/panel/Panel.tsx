// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

import React from "react";
import classes from "./panel.module.css";

export default function Panel(props: { children: React.ReactNode }) {
	return <div className={classes.panel}>{props.children}</div>;
}
