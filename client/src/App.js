import React from 'react';
import { Container } from "semantic-ui-react";
import {Quote, QuoteBuilder} from "./Quote"
import './App.css';

function App() {
  return (
    <div>
      <Container>
        <Quote />
        <QuoteBuilder />
      </Container>
    </div>
  );
}

export default App;
