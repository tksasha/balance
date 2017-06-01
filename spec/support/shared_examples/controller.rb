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

RSpec.shared_examples :new do
  before { @format ||= :js }

  describe "#new.#{ @format }" do
    before { get :new, xhr: (@format == :js), format: @format }

    it { should render_template :new }
  end
end

RSpec.shared_examples :create do
  before { @format ||= :js }

  before { expect(subject).to receive(:build_resource).and_return(resource) }

  before { expect(subject).to receive(:resource).and_return(resource) }

  describe "#create.#{ @format }" do
    context do
      before { expect(resource).to receive(:save).and_return(true) }

      before { post :create, format: @format }

      it { success.call }
    end

    context do
      before { expect(resource).to receive(:save).and_return(false) }

      before { post :create, format: @format }

      it { failure.call }
    end
  end
end

RSpec.shared_examples :destroy do
  before { @format ||= :js }

  let(:resource) { double }

  before { expect(subject).to receive(:resource).and_return(resource) }

  describe "#destroy.#{ @format }" do
    before { expect(resource).to receive(:destroy) }

    before { delete :destroy, params: { id: 1 }, format: @format }

    it { callback.call }
  end
end
