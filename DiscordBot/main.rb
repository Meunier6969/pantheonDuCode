require 'discordrb'
require 'dotenv'
Dotenv.load

bot = Discordrb::Commands::CommandBot.new token: ENV['BOT_TOKEN'], prefix: 'koyuki '

bot.mention() do |event|
	event.message.reply! "I'm alive !", mention_user: true
end

bot.command :prison, required_roles: [ENV['BIG_SISTER_ROLE_ID']]  do |event|
	event.respond "😭"
	puts 'Shutting down from Discord'
	bot.stop
end

# LET'S GO GAMBLING
bot.command(:'pile-ou-face') do |event|
	if rand(2) == 1
		event.message.reply! "Pile!", mention_user: true
	else
		event.message.reply! "Face!", mention_user: true
	end
end

at_exit { bot.stop }
bot.mode = :error

bot.run(background: true)
puts "Koyuki running, nihaha"
bot.send_message ENV['GENERAL_CHAN_ID'], 'Koyuki online!'
bot.join

# Pile ou face
# Gacha (78.5, 18.5, 3)
