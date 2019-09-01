# frozen_string_literal: true

class DateFactory
  def initialize(params = {})
    @year = params[:year].to_i

    @month = params[:month].to_i

    @day = params[:day].to_i
  end

  def build
    Date.today.change(year: year, month: month, day: day)
  rescue ArgumentError
    Date.today
  end

  private

  def year
    @year.positive? ? @year : Date.today.year
  end

  def month
    @month.positive? ? @month : Date.today.month
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
