# frozen_string_literal: true

class DateRange
  class << self
    def parse(params = {})
      begin
        Date.new(params[:year], params[:month], 1)
      rescue TypeError, Date::Error
        Time.zone.today
      end.all_month
    end
  end
end
