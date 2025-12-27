import React, { useState } from "react";

export default function UrlShortener() {
  const [longUrl, setLongUrl] = useState("");
  const [shortUrl, setShortUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [copied, setCopied] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!longUrl) return;

    setLoading(true);
    setError("");
    setShortUrl("");
    setCopied(false);

    try {
      // Create form-data
      const formData = new FormData();
      formData.append("url", longUrl);

      const response = await fetch("/api/shorten", {
        method: "POST",
        body: formData, // <-- multipart form data automatically
      });

      if (!response.ok) {
        const err = await response.json();
        throw new Error(err.fields.url);
      }

      const data = await response.json();
      setShortUrl(data.url.short_url); // depends on your response shape
    } catch (err) {
      console.log(typeof err.message)
      setError(err.message);
    } finally {
      setLoading(false);
    }
  }; const handleCopy = () => {
    navigator.clipboard.writeText(shortUrl);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  // Dummy handler for Privacy (does nothing for now)
  const handlePrivacyClick = (e) => {
    e.preventDefault();
    // TODO: implement privacy modal/page later
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900 relative overflow-hidden flex flex-col">
      {/* Animated background elements */}
      <div className="absolute inset-0 overflow-hidden">
        <div className="absolute -top-40 -right-40 w-80 h-80 bg-purple-500 rounded-full mix-blend-multiply filter blur-xl opacity-20 animate-pulse"></div>
        <div className="absolute -bottom-40 -left-40 w-80 h-80 bg-cyan-500 rounded-full mix-blend-multiply filter blur-xl opacity-20 animate-pulse delay-1000"></div>
        <div className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-80 h-80 bg-pink-500 rounded-full mix-blend-multiply filter blur-xl opacity-20 animate-pulse delay-2000"></div>
      </div>

      <div className="relative z-10 flex-1 flex flex-col items-center justify-center px-4 py-12">
        {/* Header */}
        <div className="text-center mb-12 max-w-4xl">
          <div className="inline-block mb-4 px-4 py-2 bg-gradient-to-r from-cyan-500/20 to-purple-500/20 rounded-full border border-cyan-500/30 backdrop-blur-sm">
            <span className="text-cyan-300 text-sm font-medium">
              Free • Simple • Fast
            </span>
          </div>

          <h1 className="text-5xl sm:text-6xl lg:text-7xl font-bold text-white mb-6 leading-tight">
            Shorten Links
            <br />
            <span className="bg-gradient-to-r from-cyan-400 via-purple-400 to-pink-400 text-transparent bg-clip-text">
              In Seconds
            </span>
          </h1>

          <p className="text-xl text-slate-300 mb-8 max-w-2xl mx-auto leading-relaxed">
            Clean, simple URL shortening without the hassle. No sign-up
            required, no limits, completely free.
          </p>

          {/* Simple Value Props */}
          <div className="flex flex-wrap justify-center gap-6 text-slate-300 text-sm">
            <div className="flex items-center gap-2">
              <svg
                className="w-5 h-5 text-green-400"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span>100% Free</span>
            </div>
            <div className="flex items-center gap-2">
              <svg
                className="w-5 h-5 text-green-400"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span>No Registration</span>
            </div>
            <div className="flex items-center gap-2">
              <svg
                className="w-5 h-5 text-green-400"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span>Instant Results</span>
            </div>
          </div>
        </div>

        {/* Main Card */}
        <div className="w-full max-w-2xl">
          <div className="bg-white/10 backdrop-blur-xl rounded-3xl shadow-2xl border border-white/20 p-8 sm:p-12">
            <div className="space-y-6">
              <div className="relative">
                <input
                  type="url"
                  value={longUrl}
                  onChange={(e) => setLongUrl(e.target.value)}
                  placeholder="Paste your long URL here..."
                  className="w-full px-6 py-4 bg-white/5 border border-white/20 rounded-2xl text-white placeholder-slate-400 focus:outline-none focus:ring-2 focus:ring-cyan-400 focus:border-transparent transition-all duration-300 backdrop-blur-sm text-lg"
                  required
                />
              </div>

              <button
                onClick={handleSubmit}
                disabled={loading || !longUrl}
                className="w-full bg-gradient-to-r from-cyan-500 to-purple-600 text-white font-semibold px-8 py-4 rounded-2xl hover:from-cyan-600 hover:to-purple-700 transition-all duration-300 transform hover:scale-105 disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none shadow-lg shadow-purple-500/50 text-lg"
              >
                {loading ? (
                  <span className="flex items-center justify-center gap-2">
                    <svg
                      className="animate-spin h-5 w-5"
                      viewBox="0 0 24 24"
                    >
                      <circle
                        className="opacity-25"
                        cx="12"
                        cy="12"
                        r="10"
                        stroke="currentColor"
                        strokeWidth="4"
                        fill="none"
                      ></circle>
                      <path
                        className="opacity-75"
                        fill="currentColor"
                        d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                      ></path>
                    </svg>
                    Shortening...
                  </span>
                ) : (
                  <span className="flex items-center justify-center gap-2">
                    <svg
                      className="w-5 h-5"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        strokeWidth={2}
                        d="M13 10V3L4 14h7v7l9-11h-7z"
                      />
                    </svg>
                    Shorten URL
                  </span>
                )}
              </button>
            </div>

            {/* Success Result */}
            {shortUrl && (
              <div className="mt-6 p-6 bg-gradient-to-r from-cyan-500/10 to-purple-500/10 border border-cyan-500/30 rounded-2xl backdrop-blur-sm animate-fadeIn">
                <div className="flex items-center justify-between gap-4 flex-wrap">
                  <div className="flex-1 min-w-0">
                    <div className="text-xs text-slate-400 mb-2 uppercase tracking-wide">
                      Your shortened URL
                    </div>
                    <div className="text-cyan-300 font-mono text-xl break-all">
                      {shortUrl}
                    </div>
                  </div>
                  <button
                    onClick={handleCopy}
                    className="flex-shrink-0 px-6 py-3 bg-white/10 hover:bg-white/20 text-white rounded-xl transition-all duration-200 border border-white/20 font-medium"
                  >
                    {copied ? (
                      <span className="flex items-center gap-2">
                        <svg
                          className="w-5 h-5 text-green-400"
                          fill="none"
                          stroke="currentColor"
                          viewBox="0 0 24 24"
                        >
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M5 13l4 4L19 7"
                          />
                        </svg>
                        Copied!
                      </span>
                    ) : (
                      <span className="flex items-center gap-2">
                        <svg
                          className="w-5 h-5"
                          fill="none"
                          stroke="currentColor"
                          viewBox="0 0 24 24"
                        >
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"
                          />
                        </svg>
                        Copy
                      </span>
                    )}
                  </button>
                </div>
              </div>
            )}

            {/* Error */}
            {error && (
              <div className="mt-6 p-4 bg-red-500/10 border border-red-500/30 rounded-2xl backdrop-blur-sm">
                <div className="flex items-center gap-3 text-red-300">
                  <svg
                    className="w-5 h-5 flex-shrink-0"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                    />
                  </svg>
                  <span>{error}</span>
                </div>
              </div>
            )}
          </div>

          {/* Coming Soon / Future Features */}
          <div className="mt-12">
            <div className="text-center mb-6">
              <h2 className="text-2xl font-bold text-white mb-2">
                Building the Future
              </h2>
              <p className="text-slate-400">
                Features in development to enhance your experience
              </p>
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 gap-6">
              <div className="text-center p-6 bg-white/5 backdrop-blur-sm rounded-2xl border border-white/10 relative overflow-hidden group hover:bg-white/10 transition-all">
                <div className="absolute top-2 right-2 px-2 py-1 bg-cyan-500/20 text-cyan-300 text-xs rounded-full border border-cyan-500/30">
                  Coming Soon
                </div>
                <div className="w-12 h-12 mx-auto mb-4 bg-gradient-to-br from-cyan-500 to-purple-600 rounded-xl flex items-center justify-center opacity-70 group-hover:opacity-100 transition-opacity">
                  <svg
                    className="w-6 h-6 text-white"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
                    />
                  </svg>
                </div>
                <h3 className="text-white font-semibold mb-2">
                  Click Analytics
                </h3>
                <p className="text-slate-400 text-sm">
                  Track how many people click your links
                </p>
              </div>

              <div className="text-center p-6 bg-white/5 backdrop-blur-sm rounded-2xl border border-white/10 relative overflow-hidden group hover:bg-white/10 transition-all">
                <div className="absolute top-2 right-2 px-2 py-1 bg-cyan-500/20 text-cyan-300 text-xs rounded-full border border-cyan-500/30">
                  Coming Soon
                </div>
                <div className="w-12 h-12 mx-auto mb-4 bg-gradient-to-br from-purple-500 to-pink-600 rounded-xl flex items-center justify-center opacity-70 group-hover:opacity-100 transition-opacity">
                  <svg
                    className="w-6 h-6 text-white"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01"
                    />
                  </svg>
                </div>
                <h3 className="text-white font-semibold mb-2">
                  Custom Aliases
                </h3>
                <p className="text-slate-400 text-sm">
                  Choose your own short URL names
                </p>
              </div>
            </div>
          </div>

          {/* Why This Matters Section */}
          <div className="mt-16 text-center">
            <div className="bg-gradient-to-r from-white/5 to-white/10 backdrop-blur-sm rounded-3xl p-8 border border-white/10">
              <h2 className="text-3xl font-bold text-white mb-4">
                Why It Matters
              </h2>
              <p className="text-slate-300 text-lg leading-relaxed max-w-3xl mx-auto">
                We're starting with the fundamentals—reliable, fast URL
                shortening that just works. As we grow, we'll add the
                features users actually need, not bloated extras they
                don't.
                <span className="text-cyan-300 font-medium">
                  {" "}
                  Simple today, powerful tomorrow.
                </span>
              </p>
            </div>
          </div>
        </div>
      </div>

      {/* Footer */}
      <footer className="relative z-10 border-t border-white/10 bg-black/20 backdrop-blur-xl">
        <div className="max-w-5xl mx-auto px-4 py-6 flex flex-col sm:flex-row items-center justify-between gap-3 text-sm text-slate-300">
          <div className="flex items-center gap-2">
            <span className="text-xs uppercase tracking-widest text-cyan-300/80">
              URL Shortener
            </span>
            <span className="hidden sm:inline text-slate-500">•</span>
            <span className="text-slate-400">
              Made with focus on speed and simplicity
            </span>
          </div>

          <div className="flex items-center gap-4">
            <a
              href="mailto:lokesh3721@gmail.com"
              className="text-cyan-300 hover:text-cyan-200 transition-colors"
            >
              Contact
            </a>
            <button
              onClick={handlePrivacyClick}
              className="text-slate-400 hover:text-slate-200 transition-colors"
            >
              Privacy
            </button>
          </div>
        </div>
      </footer>

      <style>{`
            @keyframes fadeIn {
               from {
                  opacity: 0;
                  transform: translateY(10px);
               }
               to {
                  opacity: 1;
                  transform: translateY(0);
               }
            }
            .animate-fadeIn {
               animation: fadeIn 0.5s ease-out;
            }
         `}</style>
    </div>
  );
}
