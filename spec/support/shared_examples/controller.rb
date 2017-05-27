RSpec.shared_examples :index do
  before { @format ||= :js }

  describe "#index.#{ @format }" do
    before { get :index, xhr: (@format == :js), format: @format }

    it { should render_template :index }
  end
end
