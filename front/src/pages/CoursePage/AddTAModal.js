import React from "react";
import { Modal, Header, Button, Icon, Input } from "semantic-ui-react";
import Autocomplete from "./AutoComplete";
import { useQuery, gql } from "@apollo/client";

// users(start: Int!=0, amount: Int!=5): [User!]!

const USERS_QUERY = gql`
  query GetUsers($start: Int!, $amount: Int!) {
    users(start: $start, amount: $amount) {
      username
    }
  }
`;

const AddTAModal = ({ open, setOpen, courseID, students }) => {
  const { data, loading, error } = useQuery(USERS_QUERY, {
    fetchPolicy: "cache-and-network",
    nextFetchPolicy: "cache-first",
    variables: {
      start: 0,
      amount: 1000
    }
  });

  let usernames;
  if (!loading) {
    usernames = students.map(user => user.username);
  }

  console.log("usernamessssss:", usernames);

  return (
    <Modal open={open} size="tiny">
      <Header icon="add user" content="Add TA" />
      <Modal.Content>
        {/* <Input icon="users" iconPosition="left" placeholder="Search users..." /> */}
        {!loading && (
          <Autocomplete suggestions={usernames} courseID={courseID} />
        )}
      </Modal.Content>
      <Modal.Actions>
        <Button color="green" onClick={() => setOpen({ addingTA: false })}>
          {/* <Icon name="checkmark" /> */}
          Add
        </Button>
        <Button color="red" onClick={() => setOpen({ addingTA: false })}>
          {/* <Icon name="remove" />  */}
          Cancel
        </Button>
      </Modal.Actions>
    </Modal>
  );
};

export default AddTAModal;
