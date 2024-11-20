/** @type {import('next').NextConfig} */
const nextConfig = {
    rewrites: async () => {
        return [
            {
              source: '/api/:path*',
              destination: 'http://api:8080/:path*',
            },
        ];
    },
};

export default nextConfig;
