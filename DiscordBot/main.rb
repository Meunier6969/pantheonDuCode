require 'discordrb'
require 'dotenv'
Dotenv.load

bot = Discordrb::Commands::CommandBot.new token: ENV['BOT_TOKEN'], prefix: 'koyuki '

# Stop the bot from Discord
bot.command(:prison, required_roles: [ENV['BIG_SISTER_ROLE_ID']]) do |event|
	event.respond "😭"
	bot.stop
end

at_exit { bot.stop }
bot.mode = :error

bot.run(background: true)
puts "Koyuki running, nihaha"
bot.send_message ENV['GENERAL_CHAN_ID'], 'Koyuki online!'
bot.join

# Gacha 
# 1* 78.5%
# 2* 18.5%
# 3* 3%
# >=2* garanteed on every 10th pull
