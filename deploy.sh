set -e

echo "🛰 Pulling latest code..."
git pull origin main

echo "🛑 Stopping existing containers..."
docker-compose down

echo "🔄 Rebuilding and starting containers..."
docker-compose up --build -d

echo "✅ Deployment complete!"
