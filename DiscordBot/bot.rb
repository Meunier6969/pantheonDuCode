require 'discordrb'
require 'dotenv'
Dotenv.load

bot = Discordrb::Commands::CommandBot.new token: ENV['BOT_TOKEN'], prefix: 'koyuki '

user_balance = {}

bot.command :a do |event| 
	"#{bot.emoji}"
end


# Admin commands
bot.command :prison, required_roles: [ENV['BIG_SISTER_ROLE_ID']], description: 'Stop the bot' do |event|
	event.respond "😭"
	puts 'Shutting down from Discord'
	bot.stop
end

bot.command :'balance-all', description: 'Reply with everyones balance' do |event|
	event.respond user_balance.to_s
	# for user in user_balance.each_key do
	# 	message < "#{user}"
	# end
end

# LET'S GO GAMBLING
bot.command :balance, description: 'Reply with users balance' do |event|
	next event.message.reply! 'You don\'t have any money :sob:' unless user_balance[event.author.username] != nil
	event.message.reply! "#{event.author.mention} has #{user_balance[event.author.username]} #{bot.emoji(1296270216203997315).use}"
end

bot.command :'add-money', description: "Add money to users balance" do |event, *args|
	user_balance[event.author.username] ||= 0 # Set value to 0 if the key doesn't exist
	user_balance[event.author.username] += args[0].to_i
end

bot.command :'pile-ou-face', description: "Simple pile ou face"  do |event|
	if rand 2 == 1
		event.message.reply! "Pile!", mention_user: true
	else
		event.message.reply! "Face!", mention_user: true
	end
end

# Hehe
bot.command :balls do |event|
	event.message.reply! "balls, cock even", mention_user: true
end

at_exit { bot.stop }
bot.mode = :error

bot.run(background: true)
puts "Koyuki running, nihaha"
bot.send_message ENV['GENERAL_CHAN_ID'], 'Koyuki online!'
bot.join

# Pile ou face
# Gacha (78.5, 18.5, 3)
