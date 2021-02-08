import React, { useState } from "react";
import { Menu, Segment } from "semantic-ui-react";
import { useHistory } from "react-router-dom";

const LoggedInHeader = props => {
  return (
    <Segment inverted style={{ borderRadius: "0px" }} fixed="top">
      <Menu inverted pointing secondary>
        {props.isMobile && (
          <Menu.Item
            icon="bars"
            onClick={() => {
              props.setSidebarIsOpen(!props.sidebarIsOpen);
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

  if (state.activeItem === "") {
    setState({ activeItem: "Homepage" });
  }

  console.log("active item:", state.activeItem);

  const history = useHistory();
  //   const auth_token = localStorage.getItem(constants.AUTH_TOKEN);

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
      // case "Logout":
      //   localStorage.removeItem(constants.AUTH_TOKEN);
      //   history.push("/");
      //   window.location.reload(false);
      //   break;
      // case "Dashboard":
      //   history.push("/dashboard");
      //   break
      // case "Homepage":
      //   history.push("/")
      //   break
    }
  }

  return (
    <div>
      {!true ? (
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
          // setToken={props.setToken}
          // isMobile={props.isMobile}
          // setSidebarIsOpen={props.setSidebarIsOpen}
          // sidebarIsOpen={props.sidebarIsOpen}
        />
      )}
    </div>
  );
};

export default Header;
