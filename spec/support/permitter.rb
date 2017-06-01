module Permitter
  def permit! params
    ActionController::Parameters.new(params).permit!
  end
end
