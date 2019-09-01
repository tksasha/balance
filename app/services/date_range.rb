# frozen_string_literal: true

class DateRange
  def initialize(date)
    @date = date
  end

  def month
    @date.beginning_of_month..@date.end_of_month
  end

  class << self
    def month(*args)
      new(*args).month
    end
  end
end
