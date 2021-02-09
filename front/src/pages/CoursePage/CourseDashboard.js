import React from "react";
import { Segment, Card, Sidebar, Label, Icon, Grid } from "semantic-ui-react";
import SideBar from "./CourseSidebar.js";
import './CourseDashboard.css'

const contents = [
  {
    title: "title1",
    timestamp: "time1",
    uploadedBY: "uploadedBy1",
    approvedBY: "approvedBy1",
    tags: ["tags1-1-sdfkjnsd", "tags1-2", "tags1-3"]
  },
  {
    title: "title2",
    timestamp: "time2",
    uploadedBY: "uploadedBy2",
    approvedBY: "approvedBy2",
    tags: ["tags2-1", "tags2-2", "tags2-3"]
  },
  {
    title: "title3",
    timestamp: "time3",
    uploadedBY: "uploadedBy3",
    approvedBY: "approvedBy3",
    tags: ["tags3-1", "tags3-2", "tags3-3"]
  },
  {
    title: "title4",
    timestamp: "time4",
    uploadedBY: "uploadedBy4",
    approvedBY: "approvedBy4",
    tags: ["tags4-1", "tags4-2", "tags4-3"]
  },
  {
    title: "title5",
    timestamp: "time5",
    uploadedBY: "uploadedBy5",
    approvedBY: "approvedBy5",
    tags: ["tags5-1", "tags5-2", "tags5-3"]
  },
  {
    title: "title6",
    timestamp: "time6",
    uploadedBY: "uploadedBy6",
    approvedBY: "approvedBy6",
    tags: ["tags6-1", "tags6-2", "tags6-3"]
  }
];

const ContentCard = ({ title, time, uploadedBY, approvedBY, tags }) => {
  return (
    <div>
      <Card fluid className="Content">
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
      </Card>
    </div>
  );
};

function CourseDashboard(props) {
  return (
    <div>
      <SideBar />
      <Segment
        style={{
          position: "absolute",
          left: props.isMobile ? 0 : 250,
          right: 0,
          margin: 30,
          top: 70
        }}
      >
        <Grid columns={2} stackable>
          {contents.map(content => {
            return (
              <Grid.Column textAlign="left">
                <ContentCard
                  title={content.title}
                  time={content.timestamp}
                  uploadedBY={content.uploadedBY}
                  approvedBY={content.approvedBY}
                  tags={content.tags}
                />
              </Grid.Column>
            );
          })}
        </Grid>
      </Segment>
    </div>
  );
}

export default CourseDashboard;
