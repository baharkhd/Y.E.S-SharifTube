import React, {useState} from "react";
import {Sidebar, Menu, Button, Icon, List} from "semantic-ui-react";
import {useHistory, Link, useParams} from "react-router-dom";
import {gql, useMutation} from "@apollo/client";
import AddTAModal from "./AddTAModal";
import PendingPage from "./PendingPage";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faSearchPlus} from "@fortawesome/free-solid-svg-icons/faSearchPlus";
import {faHeading} from "@fortawesome/free-solid-svg-icons/faHeading";
import {faChalkboardTeacher} from "@fortawesome/free-solid-svg-icons/faChalkboardTeacher";
import {faHandsHelping} from "@fortawesome/free-solid-svg-icons/faHandsHelping";
import {faUserPlus} from "@fortawesome/free-solid-svg-icons/faUserPlus";
import {faFileUpload} from "@fortawesome/free-solid-svg-icons/faFileUpload";
import {faUpload} from "@fortawesome/free-solid-svg-icons/faUpload";
import {faFileImport} from "@fortawesome/free-solid-svg-icons/faFileImport";

const sidePanelLStyle = {
    backgroundColor: '#5383ff',
    width: 250,
    top: 70
}

const sidePanelItemLStyle = {
    overflow: 'auto',
    color: 'white',
    fontSize: '20px'
}

const TAs = [
    "folan1",
    "folan2",
    "folan3",
    "folan4",
    "folan5",
    "folan6",
    "folan7"
];

// const ADD_TA_MUTATION = gql`
//   mutation AddTA($courseID: String!, $targetUsername: String!) {
//     promoteUserToTA(courseID: $courseID, targetUsername: $targetUsername) {
//       __typename
//       ... on Course {
//         id
//         title
//         summary
//         createdAt
//       }
//       ... on Exception {
//         message
//       }
//     }
//   }
// `;

function SideBar(props) {
    const [state, setState] = useState({
        activeItem: "Personal Information",
        addingTA: false
    });

    let isProfTA =
        props.courseTAs.some(ta => ta.username === props.username) || props.isProf;

    console.log(
        "{}{}{}",
        props.courseTAs.some(ta => ta.username === props.username),
        props.isProf,
        isProfTA
    );

    // const [promoteUserToTA] = useMutation(ADD_TA_MUTATION, {
    //   variables: {}
    // })

    let {id} = useParams();
    id = id.substring(1);

    const history = useHistory();

    const handleItemClick = (e, {name}) => setState({activeItem: name});

    console.log("course TAs:", props.courseTAs);
    console.log("username:", props.username);
    console.log(
        "is a TA:",
        props.courseTAs.some(ta => ta.username === props.username)
    );

    let uploadPath = isProfTA
        ? "/course:" + id + "/upload"
        : "/course:" + id + "/offer";

    return (
        <Sidebar
            as={Menu}
            animation="overlay"
            icon="labeled"
            direction="left"
            vertical
            visible={props.isMobile ? props.sidebarIsOpen : true}
            width="thin"
            style={sidePanelLStyle}
        >
            <AddTAModal
                open={state.addingTA}
                setOpen={setState}
                courseID={id}
                students={props.students}
            />
            <Menu.Item style={sidePanelItemLStyle} as="a">
                <FontAwesomeIcon size='1x'
                                 icon={faHeading}/><span><b>&nbsp;&nbsp;{props.courseTitle ? props.courseTitle : ""}</b></span>
            </Menu.Item>
            <Menu.Item style={sidePanelItemLStyle} as="a">
                <FontAwesomeIcon size='1x'
                                 icon={faChalkboardTeacher}/><span><b>&nbsp;&nbsp;{props.courseProf.name ? props.courseProf.name : ""}</b></span>
            </Menu.Item>
            <Menu.Item as="a" style={sidePanelItemLStyle}>
                <FontAwesomeIcon size='1x' icon={faHandsHelping}/><span><b>&nbsp;&nbsp;TA List</b></span>
                <List>
                    {props.courseTAs.map(TA => {
                        return (
                            <List.Item as="li">
                                {/* <List.Icon name="user" /> */}
                                {/* <List.Content>{TA}</List.Content> */}
                                {TA.name}
                            </List.Item>
                        );
                    })}
                </List>
            </Menu.Item>

            {isProfTA && (
                <Menu.Item
                    positive
                    onClick={() => {
                        setState({...state, addingTA: true});
                    }}
                    style={sidePanelItemLStyle}
                >
                    <FontAwesomeIcon size='2x' icon={faUserPlus}/><br/><br/>
                    Add TA
                </Menu.Item>
            )}
            <Link to={uploadPath + "/video"}>
                <Menu.Item style={sidePanelItemLStyle} as="a">
                    <FontAwesomeIcon size='2x' icon={faUpload}/><br/><br/>
                    Upload Videos
                </Menu.Item>
            </Link>
            <Link to={uploadPath + "/attachment"}>
                <Menu.Item style={sidePanelItemLStyle} as="a">
                    <FontAwesomeIcon size='2x' icon={faFileUpload}/><br/><br/>
                    Upload Attachments
                </Menu.Item>
            </Link>
            {isProfTA && (
                // <Link to={"/course:" + id + "/pendings"} component={PendingPage} >
                <Link to={"/course:" + id + "/pendings"}>
                    <Menu.Item style={sidePanelItemLStyle} as="a">
                        <FontAwesomeIcon size='2x' icon={faFileImport}/><br/><br/>
                        Pending Contents
                    </Menu.Item>
                </Link>
                // </Link>
            )}
        </Sidebar>
    );
}

export default SideBar;
