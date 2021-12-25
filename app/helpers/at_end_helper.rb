# frozen_string_literal: true

module AtEndHelper
  def at_end
    CalculateAtEndService.call(params[:currency])
  end
end
