import React, { useState } from "react";
import { Menu, Segment } from "semantic-ui-react";

const LoggedInHeader = props => {
  return (
    // <Menu fixed="top" inverted style={{ borderRadius: "0px" }}>
    //   {props.isMobile && (
    //     <Menu.Item
    //       icon="bars"
    //       onClick={() => {
    //         props.setSidebarIsOpen(!props.sidebarIsOpen);
    //       }}
    //     />
    //   )}
    //   <Menu.Item
    //     name="Logout"
    //     active={props.state.activeItem === "Logout"}
    //     onClick={props.handleItemClick}
    //   />
    //   <Menu.Menu position="right">
    //     <Menu.Item
    //       name="Dashboard"
    //       active={props.state.activeItem === "Dashboard"}
    //       onClick={props.handleItemClick}
    //     />
    //     <Menu.Item
    //       name="Homepage"
    //       active={props.state.activeItem === "Homepage"}
    //       onClick={props.handleItemClick}
    //     />
    //   </Menu.Menu>
    // </Menu>

    <Segment inverted style={{ borderRadius: "0px" }}>
      <Menu
        inverted
        pointing
        secondary
        fixed="top"
        style={{ borderRadius: "0px" }}
      >
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
  return (
    <Segment inverted style={{ borderRadius: "0px" }}>
      <Menu inverted pointing secondary>
        <Menu.Item
          name="Login"
          active={props.state.activeItem === "Login"}
          onClick={props.handleItemClick}
        />
        <Menu.Item
          name="Register"
          active={props.state.activeItem === "Register"}
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

  //   const history = useHistory();
  //   const auth_token = localStorage.getItem(constants.AUTH_TOKEN);

  function handleItemClick(e, { name }) {
    setState({ activeItem: name });
    // switch (name) {
    //   case "Login":
    //     history.push("/login");
    //     break;
    //   case "Register":
    //     history.push("/register");
    //     break;
    //   case "Homepage":
    //     history.push("/");
    //     break;
    //   case "Logout":
    //     localStorage.removeItem(constants.AUTH_TOKEN);
    //     history.push("/");
    //     window.location.reload(false);
    //     break;
    //   case "Dashboard":
    //     history.push("/dashboard");
    //     break
    //   case "Homepage":
    //     history.push("/")
    //     break
    // }
  }

  return (
    <div>
      {!false ? (
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
