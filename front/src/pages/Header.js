import React, { useState } from "react";
import { useHistory, useLocation } from "react-router-dom";
import { Menu, Segment } from "semantic-ui-react";

const LoggedInHeader = props => {
  return (
    <Segment
      inverted
      style={{
        borderRadius: "0px",
        top: "0px",
        position: "fixed",
        right: "0px",
        left: "0px",
        zIndex: 100,
        boxShadow: "0 8px 6px -6px black"

        
      }}
    >
      <Menu pointing secondary inverted>
        {props.isMobile && (
          <Menu.Item
            icon="bars"
            onClick={() => {
              props.setSidebarOpen(!props.sidebarOpen);
            }}
          />
        )}
        <Menu.Item
          name="Logout"
          active={props.state.activeItem === "Logout"}
          onClick={props.handleItemClick}
        />

        <Menu.Menu position="right">
          <Menu.Item
            name="Dashboard"
            active={props.state.activeItem === "Dashboard"}
            onClick={props.handleItemClick}
          />
          <Menu.Item
            name="Homepage"
            active={props.state.activeItem === "Homepage"}
            onClick={props.handleItemClick}
          />
        </Menu.Menu>
      </Menu>
    </Segment>
  );
};

const MainHeader = props => {
  console.log("active item in main header:", props.state.activeItem);
  return (
    <Segment inverted style={{ borderRadius: "0px" }}>
      <Menu inverted pointing secondary>
        <Menu.Item
          name="Login"
          active={props.state.activeItem === "Login"}
          onClick={props.handleItemClick}
        />
        <Menu.Item
          name="Sign Up"
          active={props.state.activeItem === "Sign Up"}
          onClick={props.handleItemClick}
        />
        <Menu.Menu position="right">
          <Menu.Item
            name="Homepage"
            active={props.state.activeItem === "Homepage"}
            onClick={props.handleItemClick}
          />
        </Menu.Menu>
      </Menu>
    </Segment>
  );
};

const Header = props => {
  const [state, setState] = useState({
    activeItem: "",
    loggedIn: false
  });

  const pathname = useLocation();

  if (state.activeItem === "") {
    setState({ activeItem: "Homepage" });
  }

  console.log("active item:", state.activeItem);

  const history = useHistory();

  async function changeToken() {
    await props.setToken(undefined);
    history.push("/login");
  }

  function handleItemClick(e, { name }) {
    setState({ activeItem: name });
    switch (name) {
      case "Login":
        history.push("/login");
        break;
      case "Sign Up":
        history.push("/signup");
        break;
      case "Homepage":
        history.push("/");
        break;
      case "Logout":
        changeToken();
        break;
      case "Dashboard":
        history.push("/dashboard");
        break;
    }
  }

  return (
    <div style={{position: "fixed", width: "100%", zIndex: 5}}>
      {!props.token ? (
        <MainHeader
          handleItemClick={handleItemClick}
          state={state}
          setState={setState}
        />
      ) : (
        <LoggedInHeader
          handleItemClick={handleItemClick}
          state={state}
          setState={setState}
          isMobile={props.isMobile}
          sidebarOpen={props.sidebarOpen}
          setSidebarOpen={props.setSidebarOpen}
        />
      )}
    </div>
  );
};

export default Header;
