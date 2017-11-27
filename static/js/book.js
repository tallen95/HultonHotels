function getRoomsByState() {
	console.log("here")
	var state = $('#room_state_selection').find(":selected").text();

	$.get("/room_data/", {state: state}, function(data) {
		for (room in data) {
			var $room_entry = $($(".room_entry").html());
			$('.room_entry_title', $room_entry).text(room.hotel_name);
    			$('.room_entry_description', $room_entry).text(room.description);						

			//appends edited room entry to show rooms div
			$("#show_rooms").append(room_entry);
		}
	});
	console.log("here 2")
}
