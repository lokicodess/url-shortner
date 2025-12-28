import React, { useEffect } from "react";
import { useParams } from "react-router-dom";

function ShortUrlRedirect() {
   const { code } = useParams();

   useEffect(() => {
      // Redirect to API which will handle the 302 redirect to original URL
      window.location.href = `https://api.clck.dev/${code}`;
   }, [code]);

   return (
      <div
         style={{
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            height: "100vh",
            fontFamily: "Arial, sans-serif",
         }}
      >
         <div style={{ textAlign: "center" }}>
            <h2>Redirecting...</h2>
            <p>Taking you to your destination</p>
         </div>
      </div>
   );
}

export default ShortUrlRedirect;
