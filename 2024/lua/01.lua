-- AoC 2024 Day 1

local input, err = io.open("../inputs/01.txt", "r")
if not input then
	print("Error opening file: " .. err)
	return
end

local col1 = {}
local col2 = {}

for line in input:lines() do
	local val1, val2 = line:match("(%d+)%s+(%d+)")

	if val1 and val2 then
		table.insert(col1, tonumber(val1))
		table.insert(col2, tonumber(val2))
	end
end

input:close()

table.sort(col1)
table.sort(col2)

for i, _ in ipairs(col1) do
	print("C1: " .. col1[i])
	print("C2: " .. col2[i])
end
