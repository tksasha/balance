RSpec.shared_examples :index do
  before { @format ||= :js }

  describe "#index.#{ @format }" do
    before { get :index, xhr: (@format == :js), format: @format }

    it { should render_template :index }
  end
end

RSpec.shared_examples :show do
  before { @format ||= :js }

  describe "#show.#{ @format }" do
    before { get :show, params: { id: 1 }, xhr: (@format == :js), format: @format }

    it { should render_template :show }
  end
end
