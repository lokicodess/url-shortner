import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import UrlShortener from "./UrlShortner";
import ShortUrlRedirect from "./ShortUrlRedirect";

function App() {
   return (
      <Router>
         <Routes>
            <Route path="/" element={<UrlShortener />} />
            <Route path="/:code" element={<ShortUrlRedirect />} />
         </Routes>
      </Router>
   );
}

export default App;
