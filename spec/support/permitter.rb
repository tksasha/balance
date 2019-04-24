module Permitter
  def acp params
    ActionController::Parameters.new params
  end

  # TODO: delme
  def permit! params
    acp(params).permit!
  end
end
