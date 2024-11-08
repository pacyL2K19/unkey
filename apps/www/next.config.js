const { withContentCollections } = require("@content-collections/next");

const securityHeaders = [
  {
    key: "X-Frame-Options",
    value: "SAMEORIGIN",
  },
];
const withBundleAnalyzer = require("@next/bundle-analyzer")({
  enabled: process.env.ANALYZE === "true",
});

/** @type {import('next').NextConfig} */
const nextConfig = {
  pageExtensions: ["tsx", "mdx", "ts", "js"],
  reactStrictMode: true,
  swcMinify: true,
  async headers() {
    return [
      {
        source: "/(.*)",
        headers: securityHeaders,
      },
    ];
  },
  async rewrites() {
    return [
      {
        source: "/changelog/:slug",
        destination: "/changelog#:slug", // Matched parameters can be used in the destination
      },
      {
        source: "/docs",
        destination: "https://unkey.mintlify.dev/docs",
        permanent: true,
      },
      {
        source: "/docs/:match*",
        destination: "https://unkey.mintlify.dev/docs/:match*",
        permanent: true,
      },
    ];
  },
  async redirects() {
    return [
      {
        source: "/discord",
        destination: "https://discord.gg/fDbezjbJbD",
        permanent: true,
      },
      {
        source: "/github",
        destination: "https://github.com/unkeyed/unkey",
        permanent: true,
      },
      {
        source: "/meet",
        destination: "https://cal.com/team/unkey",
        permanent: true,
      },
    ];
  },
};

module.exports = withContentCollections(withBundleAnalyzer(nextConfig));
