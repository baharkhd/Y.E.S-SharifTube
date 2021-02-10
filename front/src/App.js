import logo from "./logo.svg";
import "./App.css";
import React, { useState } from "react";
import { Button, Input } from "semantic-ui-react";
import { Route, Switch } from "react-router-dom";
import Header from "./pages/Header.js";
import Login from "./pages/Login/Login.js";
import Signup from "./pages/Signup/Signup.js";
import Homepage from "./pages/Homepage/Homepage.js";
import Dashboard from "./pages/dashboard/Dashboard.js";
import CourseDashboard from "./pages/CoursePage/CourseDashboard.js";
import ContentPage from "./pages/Content/ContentPage.js";
import { useMediaQuery } from "react-responsive";

function App() {
  const isMobile = useMediaQuery({
    query: "(max-device-width: 570px)"
  });

  const [sidebarOpen, setSidebarOpen] = useState(false);

  return (
    <div className="App">
      <Header
        isMobile={isMobile}
        sidebarOpen={sidebarOpen}
        setSidebarOpen={setSidebarOpen}
      />
      <Switch>
        <Route exact path="/dashboard">
          <Dashboard isMobile={isMobile} sidebarOpen={sidebarOpen} />
        </Route>
        <Route exact path="/content">
          <ContentPage />
        </Route>
        {/* Todo: remove this part! */}
        <Route exact path="/course">
          <CourseDashboard />
        </Route>
        <Route exact path="/">
          <Homepage />
        </Route>
        <Route exact path="/login">
          <Login />
        </Route>
        <Route exact path="/signup">
          <Signup />
        </Route>
      </Switch>
    </div>
  );
}

export default App;
