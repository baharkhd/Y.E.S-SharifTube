import { gql, useQuery } from "@apollo/client";
import React from "react";
import { Button, Header, Icon, Modal } from "semantic-ui-react";
import Autocomplete from "./AutoComplete";

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

  return (
    <Modal open={open} size="tiny">
      <Header icon="add user" content="Add TA" />
      <Modal.Content>
        {!loading && (
          <Autocomplete suggestions={usernames} courseID={courseID} />
        )}
      </Modal.Content>
      <Modal.Actions>

        <Button color="red" onClick={() => setOpen({ addingTA: false })}>
          <Icon name="remove" /> 
          Close
        </Button>
      </Modal.Actions>
    </Modal>
  );
};

export default AddTAModal;
