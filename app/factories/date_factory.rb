# frozen_string_literal: true

class DateFactory
  def initialize(params = {})
    @year = params[:year].to_i

    @month = params[:month].to_i

    @day = params[:day].to_i
  end

  def build
    Time.zone.today.change(year: year, month: month, day: day)
  rescue ArgumentError
    Time.zone.today
  end

  private

  def year
    @year.positive? ? @year : Time.zone.today.year
  end

  def month
    @month.positive? ? @month : Time.zone.today.month
  end

  def day
    @day.positive? ? @day : 1
  end

  class << self
    def build(*args)
      new(*args).build
    end
  end
end
