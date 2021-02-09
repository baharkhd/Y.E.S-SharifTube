import logo from "./logo.svg";
import "./App.css";
import { Route, Switch } from "react-router-dom";
import Header from "./pages/Header.js";
import Login from "./pages/Login/Login.js";
import Signup from "./pages/Signup/Signup.js";
import Homepage from "./pages/Homepage/Homepage.js";
import Dashboard from "./pages/dashboard/Dashboard.js";
import CourseDashboard from './pages/CoursePage/CourseDashboard.js'

function App() {
  return (
    <div className="App">
      <Header />
      <Switch>
        <Route exact path="/dashboard">
          <Dashboard />
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
