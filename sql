select b.*, h.name as hotel_name, h.country as hotel_country, h.type as hotel_type, r.price as price from userscheme.manager m 
                        join hotelscheme.hotel h on m.hotel_id = h.id 
                        join bookingscheme.booking b on h.id = b.hotel_id 
                        join hotelscheme.room r on b.reserved_room_id = r.id
                        left join bookingscheme.booking b_pred on b.client_id = b_pred.client_id 
                where m.id = 1 AND b_pred.arrival_date < b.arrival_date;
                
                
                
                
                

select count(*) from userscheme.manager m 
                        join hotelscheme.hotel h on m.hotel_id = h.id 
                        join bookingscheme.booking b on h.id = b.hotel_id 
                        join hotelscheme.room r on b.reserved_room_id = r.id
                        left join bookingscheme.booking b_pred on b.client_id = b_pred.client_id 
                where m.id = 1 AND b_pred.arrival_date < b.arrival_date;
                
                
select b.*, h.name as hotel_name, h.country as hotel_country, h.type as hotel_type, r.price as price,
	count(*) FILTER (WHERE b_pred.is_canceled = 'true') as canceled_count,
	count(*) FILTER (WHERE b_pred.is_canceled = 'false') as not_canceled_count,
	count(*)
	from userscheme.client c 
	left join bookingscheme.booking b on c.id = b.client_id 
	left join bookingscheme.booking b_pred on c.id = b_pred.client_id
	where b_pred.arrival_date < b.arrival_date 
	group by c.id, b.id;


select b.*, 
	h.name as hotel_name, 
	h.country as hotel_country, 
	h.type as hotel_type, 
	r.price as price,
	count(*) FILTER (WHERE b_pred.is_canceled = 'true') as canceled_count,
	count(*) FILTER (WHERE b_pred.is_canceled = 'false') as not_canceled_count
		from userscheme.manager m 
			left join hotelscheme.hotel h on m.hotel_id = h.id 
			left join bookingscheme.booking b on h.id = b.hotel_id 
			left join hotelscheme.room r on b.reserved_room_id = r.id
			left join bookingscheme.booking b_pred on b.client_id = b_pred.client_id
		where m.id = 1 AND b_pred.arrival_date < b.arrival_date
	group by b.id, b.client_id, h.id, r.id;



