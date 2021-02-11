import React from "react";
import { Segment, Label, Card, Icon, Grid, Button } from "semantic-ui-react";
import { useParams } from "react-router-dom";

const pendingContents = [
  {
    title: "title1",
    timestamp: "time1",
    uploadedBY: "uploadedBy1",
    approvedBY: "approvedBy1",
    tags: ["tags1-1-sdfkjnsd", "tags1-2", "tags1-3"],
    id: "videoID1"
  },
  {
    title: "title2",
    timestamp: "time2",
    uploadedBY: "uploadedBy2",
    approvedBY: "approvedBy2",
    tags: ["tags2-1", "tags2-2", "tags2-3"],
    id: "videoID2"
  },
  {
    title: "title3",
    timestamp: "time3",
    uploadedBY: "uploadedBy3",
    approvedBY: "approvedBy3",
    tags: ["tags3-1", "tags3-2", "tags3-3"],
    id: "videoID3"
  },
  {
    title: "title4",
    timestamp: "time4",
    uploadedBY: "uploadedBy4",
    approvedBY: "approvedBy4",
    tags: ["tags4-1", "tags4-2", "tags4-3"],
    id: "videoID4"
  },
  {
    title: "title5",
    timestamp: "time5",
    uploadedBY: "uploadedBy5",
    approvedBY: "approvedBy5",
    tags: ["tags5-1", "tags5-2", "tags5-3"],
    id: "videoID5"
  },
  {
    title: "title6",
    timestamp: "time6",
    uploadedBY: "uploadedBy6",
    approvedBY: "approvedBy6",
    tags: ["tags6-1", "tags6-2", "tags6-3"],
    id: "videoID6"
  }
];

const ContentCard = ({
  title,
  time,
  uploadedBY,
  approvedBY,
  tags,
  id,
  courseID
}) => {
  return (
    <div>
      <Card fluid>
        <Card.Content>
          <Card.Header>{title}</Card.Header>
        </Card.Content>
        <Card.Content description>
          uploaded by <b>{uploadedBY}</b> and approved by <b>{approvedBY}</b> in
          time <b>{time}</b>
        </Card.Content>
        <Card.Content extra>
          {tags.map(tag => {
            return (
              <Label style={{ marginBottom: 5 }}>
                <Icon name="hashtag" /> {tag}
              </Label>
            );
          })}
        </Card.Content>

        <Card.Content>
          <Button.Group fluid>
            <Button positive>salam</Button>
            <Button color="blue">salam</Button>
            <Button color="red">salam</Button>
          </Button.Group>
        </Card.Content>
      </Card>
    </div>
  );
};

function PendingPage() {
  let { courseID } = useParams();
  courseID = courseID.substring(1);

  return (
    <Segment style={{ top: 70 }}>
      <Grid columns={3} stackable>
        {pendingContents.map(content => {
          return (
            <Grid.Column textAlign="left">
              <ContentCard
                title={content.title}
                time={content.timestamp}
                uploadedBY={content.uploadedBY}
                approvedBY={content.approvedBY}
                tags={content.tags}
                id={content.id}
                courseID={courseID}
              />
            </Grid.Column>
          );
        })}
      </Grid>
    </Segment>
  );
}

export default PendingPage;
