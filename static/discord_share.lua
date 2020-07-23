local enable = gui.new_checkbox('Enable', 'discord_rpc', false); enable:set_tooltip('Enables discord rpc')
local keycode = gui.new_keybox("Post key", "discord_rpc_keycode")
local api = gui.new_textbox("API", "discord_rpc_api_endpoint")
local apiKey = gui.new_textbox("API Key", "discord_rpc_apiKey")

api:set_value("https://your-server/api/share")
apiKey:set_value("apikey")

function on_get(data)
    if data == "OK" then chat.write("Posted") else chat.write("Could't share") end
end

function on_key_pressed(key_code, is_char)
    if enable:get_value() and key_code == tonumber(keycode:get_value()) then
        print("username: " .. info.ev0lve.username)
        print("serverIp: " .. info.game.server_ip)
        print("serverName: " .. info.game.server_name)
        print("serverMap: " .. info.game.map)
        http.post(api:get_value(), {
            username = info.ev0lve.username,
            serverIp = info.game.server_ip,
            serverName = info.game.server_name,
            serverMap = info.game.map,
            apiKey = apiKey:get_value()
        }, on_get)
    end
end