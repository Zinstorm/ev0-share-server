local enable = gui.new_checkbox('Enable', 'discord_rpc', false); enable:set_tooltip('Enables discord rpc')
local keycode = gui.new_textbox("Keycode", "discord_rpc_keycode")
keycode:set_value("75")
local api = "https://domain-server/api/share"
local apiKey = "apikey"

renderer.new_font('discord_rpc_font', 'Verdana', 25, 800, flag.new(fontflags.antialias))

function on_get(data)
    if data == "OK" then chat.write("Posted") else chat.write("Could't share") end
end

function on_key_pressed(key_code, is_char)
    if enable:get_value() and key_code == tonumber(keycode:get_value()) then
        print("username: " .. info.ev0lve.username)
        print("serverIp: " .. info.game.server_ip)
        print("serverName: " .. info.game.server_name)
        print("serverMap: " .. info.game.map)
        http.post(api, {
            username = info.ev0lve.username,
            serverIp = info.game.server_ip,
            serverName = info.game.server_name,
            serverMap = info.game.map,
            apiKey = apiKey
        }, on_get)
    end
end