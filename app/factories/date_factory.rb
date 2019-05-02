class DateFactory
  class << self
    def build params={}
      today = Date.today

      params[:year] = params[:year].to_i > 0 ? params[:year].to_i : today.year

      params[:month] = params[:month].to_i > 0 ? params[:month].to_i : today.month

      params[:day] = params[:day].to_i > 0 ? params[:day].to_i : 1

      today.change params
    end
  end
end
