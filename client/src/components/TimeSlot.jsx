const TimeSlot = (props) => 
{
    const jsonData = { "slots" : {
        "slot1": "9:00 - 9:30 AM",
        "slot2": "9:30 - 10:00 AM",
        "slot3": "10:00 - 10:30 AM",
        "slot4": "10:30 - 11:00 AM",
        "slot5": "11:00 - 11:30 AM",
        "slot6": "11:30 - 12:00 PM"
     }
    }
    const slots = jsonData.slots
    const slotsarr = Object.keys(slots).map( function(k) {
        return (  
            <button className="button">{slots[k]} </button>
        )
    });
    return (
        <div>
            {slotsarr}
        </div>
    );
};

export default TimeSlot;
